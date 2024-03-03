package context

import (
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

const AUTH_TOKEN_CONTEXT_KEY = "authToken"

// Retrieves token from auth middleware
func GrabToken(ctx *gin.Context) *auth.Token {
	tokenData, exists := ctx.Get(AUTH_TOKEN_CONTEXT_KEY)
	if !exists {
		return nil
	}
	token, _ := tokenData.(*auth.Token)
	return token
}
