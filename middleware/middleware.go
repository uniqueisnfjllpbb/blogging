package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/uniqueisnfjllpbb/blogging/utils"
	"net/http"
)

func AuthMiddleware(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	token, err := utils.ParseToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}
