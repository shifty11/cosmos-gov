package subscription

import (
	"context"
	pb "github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/subscription_service"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//goland:noinspection GoNameStartsWithPackageName
type SubscriptionServer struct {
	pb.UnimplementedSubscriptionServiceServer
	subscriptionManager *database.SubscriptionManager
}

func NewSubscriptionsServer(subscriptionManager *database.SubscriptionManager) pb.SubscriptionServiceServer {
	return &SubscriptionServer{subscriptionManager: subscriptionManager}
}

func (server *SubscriptionServer) GetSubscriptions(ctx context.Context, _ *pb.GetSubscriptionsRequest) (*pb.GetSubscriptionsResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.Internal, "invalid user")
	}

	subsDtos := server.subscriptionManager.GetSubscriptions(entUser.ChatID, entUser.Type)
	var subs []*pb.Subscription
	for _, sub := range subsDtos {
		subs = append(subs, &pb.Subscription{
			Name:         sub.Name,
			DisplayName:  sub.DisplayName,
			IsSubscribed: sub.Notify,
		})
	}
	var res = &pb.GetSubscriptionsResponse{Subscriptions: subs}
	return res, nil
}

func (server *SubscriptionServer) ToggleSubscription(ctx context.Context, req *pb.ToggleSubscriptionRequest) (*pb.ToggleSubscriptionResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.Internal, "invalid user")
	}

	isSubscribed, err := server.subscriptionManager.ToggleSubscription(entUser.ChatID, entUser.Type, req.Name)
	if err != nil {
		log.Sugar.Error("error while toggling subscription: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occured")
	}
	sub := &pb.Subscription{Name: req.Name, IsSubscribed: isSubscribed}
	var res = &pb.ToggleSubscriptionResponse{Subscription: sub}
	return res, nil
}
