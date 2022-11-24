package servers

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/minoritea/sns/rpc/db"
	"github.com/minoritea/sns/rpc/model"
	"github.com/minoritea/sns/rpc/proto/protoconnect"
	"github.com/minoritea/sns/rpc/pubsub"
)

func New() (http.Handler, error) {
	db, err := db.New(true)
	if err != nil {
		return nil, err
	}

	err = db.CreateTables(model.Post{})
	if err != nil {
		return nil, err
	}

	pubsub := pubsub.New[model.Post]()

	mux := http.NewServeMux()
	mux.Handle(protoconnect.NewMessageServiceHandler(&MessageServer{db: db, pubsub: pubsub}))
	mux.Handle(protoconnect.NewAuthenticationServiceHandler(&AuthenticationServer{db: db}))

	return mux, nil
}

func Up() error {
	handler, err := New()
	if err != nil {
		return err
	}

	return http.ListenAndServe(
		"localhost:6500",
		h2c.NewHandler(handler, &http2.Server{}),
	)
}
