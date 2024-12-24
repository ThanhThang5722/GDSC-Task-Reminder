package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(ctx *gin.Context) {
	header := ctx.Writer.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Credentials", "true")
	fmt.Println(ctx.Cookie("jwt_token"))
	if ctx.Request.Method == http.MethodOptions {
		header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		ctx.JSON(200, gin.H{"message": "OK"})
	} else {
		ctx.Next()
	}
}
