package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/oklog/ulid/v2"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minoritea/sns/rpc/db"
	"github.com/minoritea/sns/rpc/model"
	"github.com/minoritea/sns/rpc/proto"
	"github.com/minoritea/sns/rpc/proto/protoconnect"
	"github.com/minoritea/sns/rpc/pubsub"
)

type MessageServer struct {
	db     *db.Engine
	pubsub *pubsub.PubSub[model.Post]
}

var messageStream = make(chan *proto.Message)

func (s *MessageServer) OpenStream(ctx context.Context, req *connect.Request[emptypb.Empty], ss *connect.ServerStream[proto.Response]) error {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	ch, unsubscribe := s.pubsub.Subscribe()
	defer unsubscribe()

	var posts []model.Post
	err = s.db.Limit(10).Desc("id").Find(&posts)
	if err != nil {
		return err
	}

	for _, post := range posts {
		err = ss.Send(&proto.Response{
			Message: &proto.Message{Body: post.Body},
		})
		if err != nil {
			return err
		}
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case post := <-ch:
			err = ss.Send(&proto.Response{
				Message: &proto.Message{Body: post.Body},
			})
			if err != nil {
				return err
			}
		}
	}
}

func (s *MessageServer) Post(ctx context.Context, req *connect.Request[proto.Message]) (*connect.Response[emptypb.Empty], error) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()
	post := model.Post{
		ID:   ulid.Make(),
		Body: req.Msg.Body,
	}

	_, err = s.db.Insert(&post)
	if err != nil {
		return nil, err
	}

	s.pubsub.Publish(post)

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func run() error {
	db, err := db.New(true)
	if err != nil {
		return err
	}

	err = db.CreateTables(model.Post{})
	if err != nil {
		return err
	}

	pubsub := pubsub.New[model.Post]()

	server := MessageServer{db: db, pubsub: pubsub}
	_, handler := protoconnect.NewMessageServiceHandler(&server)
	err = http.ListenAndServe(
		"localhost:6500",
		h2c.NewHandler(handler, &http2.Server{}),
	)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Println(err)
	}
}
