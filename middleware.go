package cotter

import (
	"context"
	"net/http"
	"strings"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

const keySetURL = "https://www.cotter.app/api/v0/token/jwks"

// NewMiddleware creates new Cotter authentication middleware.
func NewMiddleware(ctx context.Context, apiKeyID string, opts ...Option) (func(http.Handler) http.Handler, error) {
	options := options{errorHandler: defaultErrorHandler}
	for _, opt := range opts {
		opt(&options)
	}

	autoRefresh := jwk.NewAutoRefresh(ctx)
	autoRefresh.Configure(keySetURL)

	_, err := autoRefresh.Refresh(ctx, keySetURL)
	if err != nil {
		return nil, err
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			keySet, err := autoRefresh.Fetch(r.Context(), keySetURL)
			if err != nil {
				options.errorHandler(w, r, err)
				return
			}

			userID, err := userID(keySet, r, apiKeyID)
			if err != nil {
				options.errorHandler(w, r, err)
				return
			}

			next.ServeHTTP(w, r.WithContext(SetUserID(r.Context(), userID)))
		})
	}, nil
}

func userID(keySet jwk.Set, r *http.Request, apiKeyID string) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", unauthorized("authorization header is missing")
	}

	splitToken := strings.Fields(authHeader)
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return "", unauthorized("invalid authorization header")
	}

	token, err := jwt.Parse([]byte(splitToken[1]),
		jwt.WithKeySet(keySet),
		jwt.WithAudience(apiKeyID),
		jwt.WithClaimValue("scope", "access"),
		jwt.WithIssuer("https://www.cotter.app"),
		jwt.WithValidate(true),
	)
	if err != nil {
		return "", unauthorized("invalid token")
	}

	return token.Subject(), nil
}
