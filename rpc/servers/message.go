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

type MessageServer struct {
	db     *db.Engine
	pubsub *pubsub.PubSub[model.UserMessage]
}

func convertUserMessageToResponse(userMessage model.UserMessage) *proto.Response {
	return &proto.Response{
		Message: &proto.Message{
			UserName: userMessage.User.Name,
			Body:     userMessage.Message.Body,
		},
	}
}

func (s *MessageServer) OpenStream(ctx context.Context, req *connect.Request[emptypb.Empty], ss *connect.ServerStream[proto.Response]) error {
	ch, unsubscribe := s.pubsub.Subscribe()
	defer unsubscribe()

	var userMessages []model.UserMessage
	err := s.db.Table("user").Join("inner", "message", "user.id = message.user_id").Limit(10).Desc("message.id").Find(&userMessages)
	if err != nil {
		return err
	}

	for i := len(userMessages); i >= 1; i-- {
		err := ss.Send(convertUserMessageToResponse(userMessages[i-1]))
		if err != nil {
			return err
		}
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case userMessage := <-ch:
			err := ss.Send(convertUserMessageToResponse(userMessage))
			if err != nil {
				return err
			}
		}
	}
}

func (s *MessageServer) Post(ctx context.Context, req *connect.Request[proto.Message]) (*connect.Response[emptypb.Empty], error) {
	user := util.GetSessionUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("session user is not found")
	}

	message := model.Message{
		ID:     model.NewID(),
		UserID: user.ID,
		Body:   req.Msg.Body,
	}

	_, err := s.db.Insert(&message)
	if err != nil {
		return nil, err
	}

	s.pubsub.Publish(model.UserMessage{User: user, Message: &message})

	return connect.NewResponse(&emptypb.Empty{}), nil
}
