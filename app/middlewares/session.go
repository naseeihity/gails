package middlewares

import (
	"gails/app/pkg/config"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

//RedisSession session middleware stored by redis init
func RedisSession() gin.HandlerFunc {
	store, err := redis.NewStore(config.RedisCfg.MaxIdle,
		config.RedisCfg.Type,
		config.RedisCfg.Host,
		config.RedisCfg.Password,
		[]byte(config.RedisCfg.Secret))
	if err != nil {
		log.Fatalf("Init redis seesion middleware failed: %v", err)
	}

	store.Options(sessions.Options{
		// for https
		// HttpOnly: true,
		MaxAge: int(60 * 60),
	})

	return sessions.Sessions("gails-session", store)
}
