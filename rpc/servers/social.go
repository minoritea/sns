package servers

import (
	"context"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minoritea/sns/rpc/db"
	"github.com/minoritea/sns/rpc/proto"
	"github.com/minoritea/sns/rpc/util"
)

type SocialServer struct {
	db *db.Engine
}

func (s *SocialServer) Follow(ctx context.Context, req *connect.Request[proto.FollowRequest]) (*connect.Response[emptypb.Empty], error) {
	user := util.GetSessionUser(ctx)
	followee, err := db.FindUserByName(s.db, req.Msg.FollowerName)
	if err != nil {
		return nil, err
	}

	err = db.FollowUser(s.db, user.ID, followee.ID)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&emptypb.Empty{})
	return res, nil
}
