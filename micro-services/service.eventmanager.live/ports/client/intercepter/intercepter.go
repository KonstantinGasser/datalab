package intercepter

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	isserService  = "service.eventmanager.live"
	expTime       = time.Second * 5
	secretService = "super_secure"
)

func WithUnary(f grpc.UnaryClientInterceptor) grpc.DialOption {
	return grpc.WithUnaryInterceptor(f)
}

// WithAuth serves as an grpc client intercepter appending the grpc context with the
// JWT user token for authentication on service level
func WithServiceAuth(ctx context.Context, method string, req interface{}, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
) error {
	accessToken, err := generateServiceToken()
	if err != nil {
		return err
	}
	meta := metadata.Pairs("datalab-service-token", accessToken)

	newCtx := metadata.NewOutgoingContext(ctx, meta)
	return invoker(newCtx, method, req, reply, cc, opts...)
}

func generateServiceToken() (string, error) {
	claims := jwt.MapClaims{
		"iat": isserService,
		"exp": time.Now().Add(expTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(secretService))
	if err != nil {
		return "", fmt.Errorf("[jwts.IssueUser] could not sign token: %v", err)
	}

	return accessToken, nil
}
