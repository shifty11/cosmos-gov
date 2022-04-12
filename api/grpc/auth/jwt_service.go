package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
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

func AccessibleRoles() map[string][]Role {
	const path = "/cosmosgov_grpc"
	const authService = path + ".AuthService/"
	const subsService = path + ".SubscriptionService/"

	return map[string][]Role{
		authService + "TokenLogin":         {Unautheticated, User},
		authService + "RefreshAccessToken": {Unautheticated, User},
		subsService + "GetSubscriptions":   {User},
		subsService + "ToggleSubscription": {User},
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

func NewJWTManager(jwtSecretKey []byte, accessTokenDuration time.Duration, refreshTokenDuration time.Duration) *JWTManager {
	return &JWTManager{
		jwtSecretKey:         jwtSecretKey,
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
		Role: User,
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