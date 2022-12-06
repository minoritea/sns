package servers

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/minoritea/sns/rpc/db"
	"github.com/minoritea/sns/rpc/interceptors"
	"github.com/minoritea/sns/rpc/model"
	"github.com/minoritea/sns/rpc/proto/protoconnect"
	"github.com/minoritea/sns/rpc/pubsub"
)

func New() (http.Handler, error) {
	db, err := db.New(true)
	if err != nil {
		return nil, err
	}

	err = db.Sync2(model.Post{}, model.User{})
	if err != nil {
		return nil, err
	}

	pubsub := pubsub.New[model.UserPost, model.UserID]()

	mux := http.NewServeMux()
	mux.Handle(protoconnect.NewPostServiceHandler(
		&PostServer{db: db, pubsub: pubsub},
		connect.WithInterceptors(
			interceptors.NewErrorLoggingInterceptor(),
			interceptors.NewAuthenticationInterceptor(db),
		),
	))
	mux.Handle(protoconnect.NewAuthenticationServiceHandler(
		&AuthenticationServer{db: db},
		connect.WithInterceptors(
			interceptors.NewErrorLoggingInterceptor(),
		),
	))

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
