package grpc

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/shifty11/cosmos-gov/ent"
	"time"
)

var jwtKey = []byte("B8D4NdpsxLaYHyeU6E7j")

//type Credentials struct {
//	Token  string `json:"token"`
//	ChatId string `json:"chat_id"`
//}

type Claims struct {
	jwt.StandardClaims
	ChatId int64  `json:"chat_id"`
	Role   string `json:"role"`
}

type JWTManager struct {
	secretKey     []byte
	tokenDuration time.Duration
}

func NewJWTManager(tokenDuration time.Duration) *JWTManager {
	return &JWTManager{jwtKey, tokenDuration}
}

func (manager *JWTManager) Generate(entUser *ent.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		ChatId: entUser.ChatID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(manager.secretKey)
}

func (manager *JWTManager) Verify(accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return manager.secretKey, nil
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
