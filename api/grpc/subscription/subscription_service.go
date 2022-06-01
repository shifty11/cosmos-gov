package subscription

import (
	"context"
	pb "github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/subscription_service"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//goland:noinspection GoNameStartsWithPackageName
type SubscriptionServer struct {
	pb.UnimplementedSubscriptionServiceServer
	subscriptionManager *database.SubscriptionManager
}

func NewSubscriptionsServer(subscriptionManager *database.SubscriptionManager) pb.SubscriptionServiceServer {
	return &SubscriptionServer{subscriptionManager: subscriptionManager}
}

func convertSubscriptionToProtobuf(entUser *ent.User, subscriptions []*database.ChatRoom) []*pb.ChatRoom {
	var rooms []*pb.ChatRoom
	for _, chatRoom := range subscriptions {
		var subs []*pb.Subscription
		for _, sub := range chatRoom.Subscriptions {
			subs = append(subs, &pb.Subscription{
				Name:         sub.Name,
				DisplayName:  sub.DisplayName,
				IsSubscribed: sub.Notify,
			})
		}
		roomType := pb.ChatRoom_TELEGRAM
		if entUser.Type == user.TypeDiscord {
			roomType = pb.ChatRoom_DISCORD
		}
		rooms = append(rooms, &pb.ChatRoom{
			Id:                  chatRoom.Id,
			Name:                chatRoom.Name,
			TYPE:                roomType,
			Subscriptions:       subs,
			WantsDraftProposals: chatRoom.WantsDraftProposals,
		})
	}
	return rooms
}

func (server *SubscriptionServer) GetSubscriptions(ctx context.Context, _ *emptypb.Empty) (*pb.GetSubscriptionsResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	log.Sugar.Debugf("GetSubscriptions for user %v (%v)", entUser.Name, entUser.UserID)

	subs := server.subscriptionManager.GetSubscriptions(entUser)
	chatRooms := convertSubscriptionToProtobuf(entUser, subs)

	var res = &pb.GetSubscriptionsResponse{ChatRooms: chatRooms}
	return res, nil
}

func (server *SubscriptionServer) ToggleSubscription(ctx context.Context, req *pb.ToggleSubscriptionRequest) (*pb.ToggleSubscriptionResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	log.Sugar.Debugf("ToggleSubscription on %v for user %v (%v)", req.Name, entUser.Name, entUser.UserID)

	isSubscribed, err := server.subscriptionManager.ToggleSubscription(entUser, req.ChatRoomId, req.Name)
	if err != nil {
		log.Sugar.Errorf("error while toggling subscription: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}
	var res = &pb.ToggleSubscriptionResponse{IsSubscribed: isSubscribed}
	return res, nil
}

func (server *SubscriptionServer) UpdateSettings(ctx context.Context, req *pb.UpdateSettingsRequest) (*pb.SettingsResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	err := server.subscriptionManager.UpdateWantsDraftProposal(entUser, req.ChatRoomId, req.WantsDraftProposals)
	if err != nil {
		log.Sugar.Errorf("error while updating WantsDraftProposals: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}
	return &pb.SettingsResponse{
		ChatRoomId:          req.ChatRoomId,
		WantsDraftProposals: req.WantsDraftProposals,
	}, nil
}
