package interceptors

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
			return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated: %v", err)
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
			return status.Errorf(codes.Unauthenticated, "Unauthenticated: %v", err)
		}
		return next(ctx, conn)
	})
}

func setUserFromCookie(ctx context.Context, engine *db.Engine, header http.Header) (context.Context, error) {
	id := util.GetSessionID(header)
	if id == "" {
		return nil, fmt.Errorf("session id is not found")
	}

	user, err := db.FindUserBySessionID(engine, model.SessionID(id))
	if err != nil {
		return nil, err
	}

	ctx = util.WithSessionUser(ctx, *user)
	return ctx, nil
}
