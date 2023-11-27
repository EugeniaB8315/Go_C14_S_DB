package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//trae el token de postman
		tokenHeader := ctx.GetHeader("tokenPostman")
		// trae token de .env variable de entorno
		tokenEnv := os.Getenv("TOKEN")

		if tokenHeader == "" || tokenHeader != tokenEnv {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token invalido",
			})
			return
		} else {
			ctx.Next()
		}

	}
}
