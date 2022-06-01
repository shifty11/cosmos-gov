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
	"time"
)

type Config struct {
	Port                 string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	JwtSecretKey         string
	TelegramToken        string
}

//goland:noinspection GoNameStartsWithPackageName
type GRPCManager struct {
	managers    *database.DbManagers
	authzClient *authz.AuthzClient
	config      *Config
}

func NewGRPCManager(managers *database.DbManagers, authzClient *authz.AuthzClient, config *Config) *GRPCManager {
	return &GRPCManager{managers: managers, authzClient: authzClient, config: config}
}

func (m GRPCManager) Start() {
	jwtManager := auth.NewJWTManager([]byte(m.config.JwtSecretKey), m.config.AccessTokenDuration, m.config.RefreshTokenDuration)
	interceptor := auth.NewAuthInterceptor(jwtManager, m.managers.UserManager, auth.AccessibleRoles())

	authServer := auth.NewAuthServer(m.managers.UserManager, jwtManager, m.config.TelegramToken)
	subscriptionServer := subscription.NewSubscriptionsServer(m.managers.SubscriptionManager)
	votePermissionServer := vote_permission.NewVotePermissionsServer(m.authzClient, m.managers.ChainManager, m.managers.WalletManager)
	adminServer := admin.NewAdminServer(m.managers.ChainManager)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	auth_service.RegisterAuthServiceServer(server, authServer)
	subscription_service.RegisterSubscriptionServiceServer(server, subscriptionServer)
	vote_permission_service.RegisterVotePermissionServiceServer(server, votePermissionServer)
	admin_service.RegisterAdminServiceServer(server, adminServer)

	lis, err := net.Listen("tcp", m.config.Port)
	if err != nil {
		log.Sugar.Fatalf("failed to listen: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Sugar.Fatalf("failed to serve grpc: %v", err)
	}
}
