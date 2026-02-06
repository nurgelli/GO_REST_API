package middlewares

import (
	"GO_REST_API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticator(context *gin.Context){
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not authorized"})
		return 
	}

	context.Set("userId", userId)
	context.Next()

}