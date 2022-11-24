package servers

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minoritea/sns/rpc/db"
	"github.com/minoritea/sns/rpc/model"
	"github.com/minoritea/sns/rpc/proto"
)

type AuthenticationServer struct {
	db *db.Engine
}

func (a *AuthenticationServer) SignUp(ctx context.Context, req *connect.Request[proto.SignUpRequest]) (*connect.Response[emptypb.Empty], error) {
	_, err := a.db.Insert(&model.User{ID: ulid.Make(), Name: req.Msg.Name, Password: []byte(req.Msg.Password)})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (a *AuthenticationServer) SignIn(ctx context.Context, req *connect.Request[proto.SignInRequest]) (*connect.Response[emptypb.Empty], error) {
	_, err := a.db.Where("name = ? and password = ?", req.Msg.Name, []byte(req.Msg.Password)).Get(&model.User{})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (a *AuthenticationServer) IsSginedIn(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	return connect.NewResponse(&emptypb.Empty{}), nil
}
