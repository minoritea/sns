package servers

import (
	"context"
	"fmt"

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
	session, err := db.Transaction(a.db, func(_db db.DB) (*model.Session, error) {
		user, err := db.CreateUser(_db, req.Msg)
		if err != nil {
			return nil, err
		}
		session, err := db.CreateSession(_db, user.ID)
		if err != nil {
			return nil, err
		}
		return session, nil
	})
	if err != nil {
		return nil, err
	}

	cookie := util.CreateSessionCookie(session.ID.String())

	res := connect.NewResponse(&emptypb.Empty{})
	res.Header().Set("set-cookie", cookie.String())
	return res, nil
}

func (a *AuthenticationServer) SignIn(ctx context.Context, req *connect.Request[proto.SignInRequest]) (*connect.Response[emptypb.Empty], error) {
	user, err := db.FindUserByName(a.db, req.Msg.GetName())
	if err != nil {
		return nil, err
	}

	err = util.Authenticate(user, req.Msg.GetPassword())
	if err != nil {
		return nil, err
	}

	session, err := db.CreateSession(a.db, user.ID)
	if err != nil {
		return nil, err
	}

	cookie := util.CreateSessionCookie(session.ID.String())

	res := connect.NewResponse(&emptypb.Empty{})
	res.Header().Set("set-cookie", cookie.String())
	return res, nil
}

func (a *AuthenticationServer) IsSignedIn(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	id := util.GetSessionID(req.Header())
	if id == "" {
		return nil, fmt.Errorf("session id is not found")
	}

	_, err := db.FindUserBySessionID(a.db, model.SessionID(id))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
