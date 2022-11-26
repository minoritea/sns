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
	"github.com/minoritea/sns/rpc/util"
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

func (a *AuthenticationServer) IsSignedIn(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	id := util.GetSessionID(req.Header())
	if id == "" {
		return nil, fmt.Errorf("session id is not found")
	}

	_, err := db.GetUser(a.db, id)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
