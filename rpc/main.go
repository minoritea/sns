package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minoritea/sns/build/proto"
	"github.com/minoritea/sns/build/proto/protoconnect"
)

type MessageServer struct{}

var messageStream = make(chan *proto.Message)

func (s *MessageServer) OpenStream(ctx context.Context, req *connect.Request[emptypb.Empty], ss *connect.ServerStream[proto.Response]) error {
	for {
		select {
		case <-ctx.Done():
			log.Println("stream is finished")
			return nil
		case msg := <-messageStream:
			ss.Send(&proto.Response{
				Message: msg,
			})
		}
	}
}

func (s *MessageServer) Post(ctx context.Context, req *connect.Request[proto.Message]) (*connect.Response[emptypb.Empty], error) {
	messageStream <- req.Msg
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func main() {
	var server MessageServer
	_, handler := protoconnect.NewMessageServiceHandler(&server)
	err := http.ListenAndServe(
		"localhost:6500",
		h2c.NewHandler(handler, &http2.Server{}),
	)
	if err != nil {
		log.Println(err)
	}
}
