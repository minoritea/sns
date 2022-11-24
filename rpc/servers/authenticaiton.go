package servers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minoritea/sns/rpc/db"
	"github.com/minoritea/sns/rpc/model"
	"github.com/minoritea/sns/rpc/proto"
)

type AuthenticationServer struct {
	db *db.Engine
}

func (a *AuthenticationServer) SignUp(ctx context.Context, req *connect.Request[proto.SignUpRequest]) (*connect.Response[emptypb.Empty], error) {
	user := model.User{ID: model.NewID(), Name: req.Msg.Name, Password: req.Msg.Password}
	_, err := a.db.Insert(&user)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&emptypb.Empty{})
	cookie := http.Cookie{
		Name:     "id",
		Value:    user.ID.String(),
		Path:     "/rpc",
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	}
	res.Header().Set("set-cookie", cookie.String())
	return res, nil
}

func (a *AuthenticationServer) SignIn(ctx context.Context, req *connect.Request[proto.SignInRequest]) (*connect.Response[emptypb.Empty], error) {
	has, err := a.db.Where("name = ? and password = ?", req.Msg.Name, req.Msg.Password).Get(&model.User{})
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("user is not found")
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (a *AuthenticationServer) IsSignedIn(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	hr := http.Request{Header: req.Header()}
	var id string
	for _, cookie := range hr.Cookies() {
		if cookie.Name == "id" {
			id = cookie.Value
			break
		}
	}

	var user model.User
	has, err := a.db.ID(id).Get(&user)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("user is not found")
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}
