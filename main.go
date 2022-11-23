package main

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minoritea/sns/build/proto"
	"github.com/minoritea/sns/build/proto/protoconnect"
)

type MessageStreamServer struct{}

func (s *MessageStreamServer) Open(ctx context.Context, req *connect.Request[emptypb.Empty], ss *connect.ServerStream[proto.Response]) error {
	messages := []*proto.Message{
		{
			Body: "Hello, this is the first message.",
		},
		{
			Body: "Hello, this is the second message.",
		},
		{
			Body: "Hello, this is the third message.",
		},
	}
	var counter int
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			switch counter {
			case 0:
				counter++
			case 1, 2, 3:
				ss.Send(&proto.Response{
					Messages: messages[counter-1 : counter],
				})
				counter++
			}
		}
	}
}

var IsDevelopment = true

func main() {
	var server MessageStreamServer
	_, handler := protoconnect.NewMessageStreamHandler(&server)
	http.ListenAndServe(
		"localhost:6500",
		h2c.NewHandler(handler, &http2.Server{}),
	)
}
