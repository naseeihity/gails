package middlewares

import (
	"gails/app/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//CheckLogin 检查登陆状态
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isSignIn := c.GetBool("state_isUserSignIn"); isSignIn {
			c.Next()
			return
		}
		// 未登录重定向到首页
		c.Redirect(http.StatusFound, "/")
		c.Abort()
	}
}

//AddState add user state
func AddState() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// 从session中获取user_id用来判断用户信息
		userID := session.Get("user_id")
		if userID != nil {
			// interface{}.(uint)会报错
			curUser, err := models.GetUserByID(userID)

			// 通过context传递用户信息参数
			if err == nil {
				c.Set("state_isUserSignIn", true)
			} else {
				c.Set("state_isUserSignIn", false)
			}
			c.Set("state_curUser", curUser)
		} else {
			c.Set("state_isUserSignIn", false)
			c.Set("state_curUser", nil)
		}

		c.Next()
	}
}
