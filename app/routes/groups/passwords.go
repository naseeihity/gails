package groups

import (
	"gails/app/controllers/passwords"
	"gails/app/middlewares"

	"github.com/gin-gonic/gin"
)

//InitPasswords :
func InitPasswords(r *gin.Engine) {
	p := r.Group("/passwords")

	// 中间件检查登陆状态
	p.Use(middlewares.CheckLogin())
	{
		p.GET("/edit", passwords.Edit)
		p.PUT("/", passwords.Update)
	}

}
