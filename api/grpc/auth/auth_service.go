package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	pb "github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/auth_service"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"time"
)

//goland:noinspection GoNameStartsWithPackageName
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userManager *database.UserManager
	jwtManager  *JWTManager
}

func NewAuthServer(userManager *database.UserManager, jwtManager *JWTManager) pb.AuthServiceServer {
	return &AuthServer{userManager: userManager, jwtManager: jwtManager}
}

func (server *AuthServer) TelegramLogin(_ context.Context, req *pb.TelegramLoginRequest) (*pb.LoginResponse, error) {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")

	s := sha256.New()
	s.Write([]byte(telegramToken))
	secretKey := s.Sum(nil)

	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(req.DataStr))
	hh := h.Sum(nil)
	hash := hex.EncodeToString(hh)
	if hash != req.Hash {
		return nil, status.Errorf(codes.Unauthenticated, "telegram data invalid")
	}

	if time.Now().Sub(time.Unix(req.AuthDate, 0)) > time.Hour {
		return nil, status.Errorf(codes.Unauthenticated, "telegram login expired")
	}

	entUser, err := server.userManager.Get(req.UserId, user.TypeTelegram)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find user: %v", err)
	}

	accessToken, err := server.jwtManager.GenerateToken(entUser, AccessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate accessToken: %v", err)
	}

	refreshToken, err := server.jwtManager.GenerateToken(entUser, RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate refreshToken: %v", err)
	}

	res := &pb.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	return res, nil
}

func (server *AuthServer) TokenLogin(_ context.Context, req *pb.TokenLoginRequest) (*pb.LoginResponse, error) {
	var userType = user.TypeTelegram
	if req.TYPE == pb.TokenLoginRequest_DISCORD {
		userType = user.TypeDiscord
	}

	entUser, err := server.userManager.ByToken(req.ChatId, userType, req.Token)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find user: %v", err)
	}

	err = server.userManager.GenerateNewLoginToken(req.ChatId, userType)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate new login token: %v", err)
	}

	accessToken, err := server.jwtManager.GenerateToken(entUser, AccessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate accessToken: %v", err)
	}

	refreshToken, err := server.jwtManager.GenerateToken(entUser, RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate refreshToken: %v", err)
	}

	res := &pb.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	return res, nil
}

func (server *AuthServer) RefreshAccessToken(_ context.Context, req *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	claims, err := server.jwtManager.Verify(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "refresh token invalid: %v", err)
	}

	entUser, err := server.userManager.Get(claims.UserId, claims.Type)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot find user: %v", err)
	}

	accessToken, err := server.jwtManager.GenerateToken(entUser, AccessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate accessToken: %v", err)
	}

	res := &pb.RefreshAccessTokenResponse{AccessToken: accessToken}
	return res, nil
}
