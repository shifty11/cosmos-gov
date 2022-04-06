package grpc

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"os"
	"time"
)

type Role string

const (
	Unautheticated Role = "Unautheticated"
	User           Role = "User"
)

type TokenType string

const (
	AccessToken  TokenType = "AccessToken"
	RefreshToken TokenType = "RefreshToken"
)

func accessibleRoles() map[string][]Role {
	const path = "/cosmosgov_grpc.AuthService/"

	return map[string][]Role{
		path + "TokenLogin":         {Unautheticated, User},
		path + "RefreshAccessToken": {Unautheticated, User},
	}
}

type Claims struct {
	jwt.StandardClaims
	ChatId int64     `json:"chat_id"`
	Type   user.Type `json:"type"`
	Role   Role      `json:"role,omitempty"`
}

type JWTManager struct {
	jwtSecretKey         []byte
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewJWTManager(accessTokenDuration time.Duration, refreshTokenDuration time.Duration) *JWTManager {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Sugar.Panic("JWT_SECRET_KEY must be set")
	}
	return &JWTManager{
		jwtSecretKey:         []byte(jwtSecretKey),
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
	}
}

func (manager *JWTManager) GenerateToken(entUser *ent.User, tokenType TokenType) (string, error) {
	expirationTime := time.Now().Add(manager.accessTokenDuration)
	if tokenType == RefreshToken {
		expirationTime = time.Now().Add(manager.refreshTokenDuration)
	}

	claims := &Claims{
		ChatId: entUser.ChatID,
		Type:   entUser.Type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(manager.jwtSecretKey)
}

func (manager *JWTManager) Verify(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return manager.jwtSecretKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
