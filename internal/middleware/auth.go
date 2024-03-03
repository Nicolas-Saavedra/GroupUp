package middleware

import (
	"github.com/gin-gonic/gin"

	"GroupUp/internal/service"
	"GroupUp/pkg/context"
)

func AuthMiddleware(authService service.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) > 8 && authHeader[:8] == "Bearer: " {
			token, err := authService.VerifyToken(authHeader[:8], ctx)
			if err == nil {
				ctx.Set(context.AUTH_TOKEN_CONTEXT_KEY, token)
			}
		}
	}
}
