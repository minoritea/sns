package main

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/go-chi/chi"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minoritea/sns/build/proto"
	"github.com/minoritea/sns/build/proto/protoconnect"
)

type MessageServer struct{}

func (s *MessageServer) List(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[proto.ListResponse], error) {
	return connect.NewResponse(
		&proto.ListResponse{
			Messages: []*proto.Message{
				{
					Body: "Hello, World",
				},
			},
		},
	), nil
}

var IsDevelopment = true

func main() {
	var server MessageServer
	router := chi.NewRouter()
	_, handler := protoconnect.NewMessageServiceHandler(&server)
	handler = http.StripPrefix("/rpc", handler)
	router.Handle("/rpc/*", handler)
	http.ListenAndServe(
		"localhost:6500",
		h2c.NewHandler(router, &http2.Server{}),
	)
}
