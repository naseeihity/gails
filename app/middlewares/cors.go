package middlewares

import (
	"gails/app/pkg/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS 跨域请求处理中间件
func CORS() gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowCredentials = true
	cfg.AllowOrigins = config.ServerCfg.AllowOrigins
	cfg.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	cfg.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}

	return cors.New(cfg)
}
