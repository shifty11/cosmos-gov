package grpc

import (
	"github.com/shifty11/cosmos-gov/api/grpc/admin"
	"github.com/shifty11/cosmos-gov/api/grpc/auth"
	_ "github.com/shifty11/cosmos-gov/api/grpc/auth"
	"github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/admin_service"
	"github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/auth_service"
	"github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/subscription_service"
	"github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/vote_permission_service"
	"github.com/shifty11/cosmos-gov/api/grpc/subscription"
	"github.com/shifty11/cosmos-gov/api/grpc/vote_permission"
	"github.com/shifty11/cosmos-gov/authz"
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

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		log.Sugar.Panic("TELEGRAM_TOKEN must be set")
	}

	managers := database.NewDefaultDbManagers()

	jwtManager := auth.NewJWTManager([]byte(jwtSecretKey), accessTokenDuration, refreshTokenDuration)
	interceptor := auth.NewAuthInterceptor(jwtManager, managers.UserManager, auth.AccessibleRoles())

	authzClient := authz.NewAuthzClient(managers.ChainManager, managers.WalletManager)

	authServer := auth.NewAuthServer(managers.UserManager, jwtManager, telegramToken)
	subscriptionServer := subscription.NewSubscriptionsServer(managers.SubscriptionManager)
	votePermissionServer := vote_permission.NewVotePermissionsServer(authzClient, managers.ChainManager, managers.WalletManager)
	adminServer := admin.NewAdminServer(managers.ChainManager)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	auth_service.RegisterAuthServiceServer(server, authServer)
	subscription_service.RegisterSubscriptionServiceServer(server, subscriptionServer)
	vote_permission_service.RegisterVotePermissionServiceServer(server, votePermissionServer)
	admin_service.RegisterAdminServiceServer(server, adminServer)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Sugar.Fatalf("failed to listen: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Sugar.Fatalf("failed to serve grpc: %v", err)
	}
}
