package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"golang.org/x/exp/slices"
	"os"
	"strconv"
	"strings"
	"time"
)

type Role string

const (
	Unauthenticated Role = "Unauthenticated"
	User            Role = "User"
	Admin           Role = "Admin"
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
	const votePermissionService = path + ".VotePermissionService/"
	const chainService = path + ".AdminService/"

	return map[string][]Role{
		authService + "TelegramLogin":                   {Unauthenticated, User, Admin},
		authService + "TokenLogin":                      {Unauthenticated, User, Admin},
		authService + "RefreshAccessToken":              {Unauthenticated, User, Admin},
		subsService + "GetSubscriptions":                {User, Admin},
		subsService + "ToggleSubscription":              {User, Admin},
		subsService + "UpdateSettings":                  {User, Admin},
		votePermissionService + "GetSupportedChains":    {User, Admin},
		votePermissionService + "RegisterWallet":        {User, Admin},
		votePermissionService + "RemoveWallet":          {User, Admin},
		votePermissionService + "GetWallets":            {User, Admin},
		votePermissionService + "RefreshVotePermission": {User, Admin},
		chainService + "GetChains":                      {Admin},
		chainService + "UpdateChain":                    {Admin},
		chainService + "ReportError":                    {Admin},
	}
}

type Claims struct {
	jwt.RegisteredClaims
	UserId int64     `json:"user_id"`
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

	admins := strings.Split(strings.Trim(os.Getenv("ADMIN_IDS"), " "), ",")
	var role = User
	if slices.Contains(admins, strconv.FormatInt(entUser.UserID, 10)) {
		role = Admin
	}

	claims := &Claims{
		UserId: entUser.UserID,
		Type:   entUser.Type,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		Role: role,
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
