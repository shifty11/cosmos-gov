package grpc

import (
	pb "github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/protobuf/auth_service"
	"github.com/shifty11/cosmos-gov/log"
	"google.golang.org/grpc"
	"net"
	"time"
)

const (
	port                 = ":50051"
	accessTokenDuration  = time.Minute * 4
	refreshTokenDuration = time.Hour * 24
)

func Start() {
	jwtManager := NewJWTManager(accessTokenDuration, refreshTokenDuration)
	authServer := NewAuthServer(jwtManager)
	interceptor := NewAuthInterceptor(jwtManager, accessibleRoles())

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	pb.RegisterAuthServiceServer(server, authServer)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Sugar.Fatalf("failed to listen: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Sugar.Fatalf("failed to serve grpc: %v", err)
	}
}
