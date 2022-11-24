package interceptors

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"

	"github.com/minoritea/sns/rpc/db"
	"github.com/minoritea/sns/rpc/model"
	"github.com/minoritea/sns/rpc/util"
)

type authenticationInterceptor struct{ db *db.Engine }

func NewAuthenticationInterceptor(db *db.Engine) connect.Interceptor {
	return &authenticationInterceptor{db: db}
}

func (a *authenticationInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		ctx, err := setUserFromCookie(ctx, a.db, req.Header())
		if err != nil {
			return nil, err
		}
		return next(ctx, req)
	})
}

func (a *authenticationInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

func (a *authenticationInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return connect.StreamingHandlerFunc(func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		ctx, err := setUserFromCookie(ctx, a.db, conn.RequestHeader())
		if err != nil {
			return err
		}
		return next(ctx, conn)
	})
}

func setUserFromCookie(ctx context.Context, db *db.Engine, header http.Header) (context.Context, error) {
	hr := http.Request{Header: header}
	var id string
	for _, cookie := range hr.Cookies() {
		if cookie.Name == "id" {
			id = cookie.Value
			break
		}
	}

	var user model.User
	has, err := db.ID(id).Get(&user)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("user is not found")
	}

	ctx = util.WithSessionUser(ctx, user)
	return ctx, nil
}
