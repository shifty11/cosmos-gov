package subscription

import (
	"context"
	pb "github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/subscription_service"
)

//goland:noinspection GoNameStartsWithPackageName
type SubscriptionServer struct {
	pb.UnimplementedSubscriptionServiceServer
}

func NewSubscriptionsServer() pb.SubscriptionServiceServer {
	return &SubscriptionServer{}
}

func (server *SubscriptionServer) GetSubscriptions(ctx context.Context, req *pb.GetSubscriptionsRequest) (*pb.GetSubscriptionsResponse, error) {
	var subs []*pb.Subscription
	var res = &pb.GetSubscriptionsResponse{Subscriptions: subs}
	return res, nil
}
