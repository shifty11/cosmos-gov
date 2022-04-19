package auth

import (
	"context"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

//goland:noinspection GoNameStartsWithPackageName
type AuthInterceptor struct {
	jwtManager      *JWTManager
	userManager     *database.UserManager
	accessibleRoles map[string][]Role
}

func NewAuthInterceptor(jwtManager *JWTManager, accessibleRoles map[string][]Role) *AuthInterceptor {
	return &AuthInterceptor{jwtManager: jwtManager, accessibleRoles: accessibleRoles}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Sugar.Debug("--> unary interceptor: ", info.FullMethod)

		ctx, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Sugar.Debug("--> stream interceptor: ", info.FullMethod)

		_, err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if slices.Contains(accessibleRoles, Unauthenticated) {
		return nil, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := strings.Replace(values[0], "Bearer ", "", 1)
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			entUser, err := interceptor.userManager.GetUser(claims.ChatId, claims.Type)
			if err != nil {
				return nil, status.Error(codes.Internal, "user not found")
			}
			return context.WithValue(ctx, "user", entUser), nil
		}
	}

	return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
