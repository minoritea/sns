package servers

import (
	"context"
	"log"

	"github.com/bufbuild/connect-go"
	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minoritea/sns/rpc/db"
	"github.com/minoritea/sns/rpc/model"
	"github.com/minoritea/sns/rpc/proto"
	"github.com/minoritea/sns/rpc/pubsub"
)

type MessageServer struct {
	db     *db.Engine
	pubsub *pubsub.PubSub[model.Post]
}

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
