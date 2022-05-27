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
	"time"
)

//goland:noinspection GoNameStartsWithPackageName
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userManager   *database.UserManager
	jwtManager    *JWTManager
	telegramToken string
}

func NewAuthServer(userManager *database.UserManager, jwtManager *JWTManager, telegramToken string) pb.AuthServiceServer {
	return &AuthServer{userManager: userManager, jwtManager: jwtManager, telegramToken: telegramToken}
}

func (server *AuthServer) secretKey1() []byte {
	s := sha256.New()
	s.Write([]byte(server.telegramToken))
	secretKey := s.Sum(nil)
	return secretKey
}

func (server *AuthServer) secretKey2() []byte {
	h1 := hmac.New(sha256.New, []byte("WebAppData"))
	h1.Write([]byte(server.telegramToken))
	secretKey := h1.Sum(nil)
	return secretKey
}

func (server *AuthServer) isValid(dataStr string, secretKey []byte, hash string) bool {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(dataStr))
	hh := h.Sum(nil)
	resultHash := hex.EncodeToString(hh)
	return resultHash == hash
}

func (server *AuthServer) TelegramLogin(_ context.Context, req *pb.TelegramLoginRequest) (*pb.LoginResponse, error) {
	if !server.isValid(req.DataStr, server.secretKey1(), req.Hash) && !server.isValid(req.DataStr, server.secretKey2(), req.Hash) {
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
