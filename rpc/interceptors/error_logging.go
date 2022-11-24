package interceptors

import (
	"context"
	"log"

	"github.com/bufbuild/connect-go"
)

func NewErrorLoggingInterceptor() connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			res, err := next(ctx, req)
			if err != nil {
				log.Printf("path: %s, error: %v", req.Spec().Procedure, err)
			}
			return res, err
		})
	})
}
