package interceptors

import (
	"context"
	"log"

	"github.com/bufbuild/connect-go"
)

type errorLoggingInterceptor struct{}

func (a *errorLoggingInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		res, err := next(ctx, req)
		if err != nil {
			log.Printf("path: %s, error: %v", req.Spec().Procedure, err)
		}
		return res, err
	})
}

func (a *errorLoggingInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

func (a *errorLoggingInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return connect.StreamingHandlerFunc(func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		err := next(ctx, conn)
		if err != nil {
			log.Printf("path: %s, error: %v", conn.Spec().Procedure, err)
		}
		return err
	})
}

func NewErrorLoggingInterceptor() connect.Interceptor {
	return &errorLoggingInterceptor{}
}
