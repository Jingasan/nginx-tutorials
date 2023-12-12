package main

// Ginのimport
import (
	"github.com/gin-gonic/gin"
)

func main() {
	// GinをReleaseモードに設定
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	// GETメソッド
	engine.GET("/api", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	// ポート番号を指定
	engine.Run(":3000")
}
