package grpc

import (
	pb "github.com/shifty11/cosmos-gov/api/grpc/cosmos-gov-grpc/go/protobuf/auth_service"
	"github.com/shifty11/cosmos-gov/log"
	"google.golang.org/grpc"
	"net"
	"time"
)

const (
	port = ":50051"
)

func accessibleRoles() map[string][]string {
	const cosmosGovPath = "/cosmosgov_grpc.CosmosGov/"

	return map[string][]string{
		//cosmosGovPath + "TokenLogin": {"admin", "user"},
	}
}

func Start() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Sugar.Fatalf("failed to listen: %v", err)
	}

	jwtManager := NewJWTManager(5 * time.Minute)
	authServer := NewAuthServer(jwtManager)
	interceptor := NewAuthInterceptor(jwtManager, accessibleRoles())

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	pb.RegisterAuthServiceServer(s, authServer)
	err = s.Serve(lis)
	if err != nil {
		log.Sugar.Fatalf("failed to serve grpc: %v", err)
	}
}
