package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/AndroDeMohawk/link-cutter/configs"
	JWT "github.com/AndroDeMohawk/link-cutter/pkg/jwt"
)

type ctxKey struct{}

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("unauthorized"))
}

func IsAuth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			writeUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		ok, data := JWT.NewJWT(config.Auth.Secret).Parse(token)
		if !ok || data == nil {
			writeUnauthed(w)
			return
		}
		ctx := context.WithValue(r.Context(), ctxKey{}, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func WithEmail(ctx context.Context) any {
	value := ctx.Value(ctxKey{})
	return value
}
