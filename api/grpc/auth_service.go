package grpc

import (
	"context"
	pb "github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/protobuf/auth_service"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userManager *database.UserManager
	jwtManager  *JWTManager
}

func NewAuthServer(jwtManager *JWTManager) pb.AuthServiceServer {
	return &AuthServer{userManager: database.NewUserManager(), jwtManager: jwtManager}
}

func (server *AuthServer) TokenLogin(_ context.Context, req *pb.TokenLoginRequest) (*pb.TokenLoginResponse, error) {
	var userType = user.TypeTelegram
	if req.TYPE == pb.TokenLoginRequest_DISCORD {
		userType = user.TypeDiscord
	}

	entUser, err := server.userManager.GetUser(req.ChatId, userType, req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	err = server.userManager.InvalidateToken(req.ChatId, userType, req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot invalidate accessToken user: %v", err)
	}

	accessToken, err := server.jwtManager.GenerateToken(entUser, AccessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate accessToken")
	}

	refreshToken, err := server.jwtManager.GenerateToken(entUser, RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate refreshToken")
	}

	res := &pb.TokenLoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	return res, nil
}

func (server *AuthServer) RefreshAccessToken(_ context.Context, req *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	claims, err := server.jwtManager.Verify(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "refresh token invalid: %v", err)
	}

	entUser, err := server.userManager.GetUser(claims.ChatId, claims.Type, "")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot find user: %v", err)
	}

	accessToken, err := server.jwtManager.GenerateToken(entUser, AccessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate accessToken")
	}

	res := &pb.RefreshAccessTokenResponse{AccessToken: accessToken}
	return res, nil
}
