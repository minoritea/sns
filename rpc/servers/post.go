package servers

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minoritea/sns/rpc/db"
	"github.com/minoritea/sns/rpc/model"
	"github.com/minoritea/sns/rpc/proto"
	"github.com/minoritea/sns/rpc/pubsub"
	"github.com/minoritea/sns/rpc/util"
)

type PostServer struct {
	db     *db.Engine
	pubsub *pubsub.PubSub[model.UserPost, model.UserID]
}

func convertUserPostToResponse(userPost model.UserPost) *proto.Response {
	return &proto.Response{
		Post: &proto.Post{
			UserName: userPost.User.Name,
			Body:     userPost.Post.Body,
		},
	}
}

func (s *PostServer) OpenStream(ctx context.Context, req *connect.Request[emptypb.Empty], ss *connect.ServerStream[proto.Response]) error {
	user := util.GetSessionUser(ctx)
	if user == nil {
		return fmt.Errorf("session user is not found")
	}
	ch, unsubscribe := s.pubsub.Subscribe(user.ID)
	defer unsubscribe()

	userPosts, err := db.FindUserPosts(s.db, 10)
	if err != nil {
		return err
	}

	for i := len(userPosts); i >= 1; i-- {
		err := ss.Send(convertUserPostToResponse(userPosts[i-1]))
		if err != nil {
			return err
		}
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case userPost := <-ch:
			err := ss.Send(convertUserPostToResponse(userPost))
			if err != nil {
				return err
			}
		}
	}
}

func (s *PostServer) Publish(ctx context.Context, req *connect.Request[proto.Post]) (*connect.Response[emptypb.Empty], error) {
	user := util.GetSessionUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("session user is not found")
	}

	post := model.Post{
		ID:     model.NewPostID(),
		UserID: user.ID,
		Body:   req.Msg.Body,
	}

	_, err := s.db.Insert(&post)
	if err != nil {
		return nil, err
	}

	s.pubsub.Publish(model.UserPost{User: user, Post: &post})

	return connect.NewResponse(&emptypb.Empty{}), nil
}
