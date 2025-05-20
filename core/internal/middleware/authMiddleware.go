package middleware

import (
	"cdisk/util"
	"context"
	"net/http"
)

type AuthMiddleware struct {
	AccessSecret string
}

func NewAuthMiddleware(Secret string) *AuthMiddleware {
	return &AuthMiddleware{Secret}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// 1. get token from header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		prefix := "Bearer "
		if len(token) < len(prefix) || token[:len(prefix)] != prefix {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token = token[len(prefix):]
		// 验证 Token
		claims, err := util.ParseToken(token, m.AccessSecret)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 假设 claims 中包含 "userId" 字段
		userId, ok := claims["userId"].(string)
		if !ok || userId == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 3. 将用户信息存入上下文
		ctx := context.WithValue(r.Context(), "userId", userId)
		next(w, r.WithContext(ctx))

		// 2. check token
		// 3. check token in redis
		// 4. check token in mysql
		// Passthrough to next handler if need
		// next(w, r)
	}
}
