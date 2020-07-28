package middlewares

import (
	"gails/app/helpers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//FlashMessage 提示消息
func FlashMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		type FlashMsg struct {
			Info []string
			Warn []string
		}

		var flash FlashMsg
		session := sessions.Default(c)
		flash.Info = helpers.InterArrToStrArr(session.Flashes("Info"))
		flash.Warn = helpers.InterArrToStrArr(session.Flashes("Warn"))
		session.Save()
		c.Set("flashMessage", flash)

		c.Next()
	}
}
