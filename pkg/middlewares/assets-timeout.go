package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CacheBuster(ctx *gin.Context) {
	urlPath := ctx.Request.URL.Path
	if strings.HasPrefix(urlPath, "/assets/css/") || strings.HasPrefix(urlPath, "/assets/js/") {
		ctx.Header("Cache-Control", "public, max-age=60, immutable")
	}
	ctx.Next()
}
