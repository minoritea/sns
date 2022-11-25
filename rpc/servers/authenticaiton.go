package servers

import (
	"context"
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
