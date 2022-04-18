package grpc

import (
	"github.com/shifty11/cosmos-gov/api/grpc/auth"
	_ "github.com/shifty11/cosmos-gov/api/grpc/auth"
	"github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/auth_service"
	"github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/subscription_service"
	"github.com/shifty11/cosmos-gov/api/grpc/subscription"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

const (
	port                 = ":50051"
	accessTokenDuration  = time.Minute * 15
	refreshTokenDuration = time.Hour * 24
)

func Start() {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Sugar.Panic("JWT_SECRET_KEY must be set")
	}
	userManager := database.NewUserManager()
	jwtManager := auth.NewJWTManager([]byte(jwtSecretKey), accessTokenDuration, refreshTokenDuration)
	interceptor := auth.NewAuthInterceptor(jwtManager, auth.AccessibleRoles())
	authServer := auth.NewAuthServer(userManager, jwtManager)
	subscriptionServer := subscription.NewSubscriptionsServer(database.NewSubscriptionManager(userManager))

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	auth_service.RegisterAuthServiceServer(server, authServer)
	subscription_service.RegisterSubscriptionServiceServer(server, subscriptionServer)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Sugar.Fatalf("failed to listen: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Sugar.Fatalf("failed to serve grpc: %v", err)
	}
}
