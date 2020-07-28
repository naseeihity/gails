package users

import (
	"gails/app/helpers"
	"gails/app/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//SignInPage 登陆页
func SignInPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signin", helpers.CommonHTMLRes("Signin", c))
}

//SignIn 登陆请求
func SignIn(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	session := sessions.Default(c)

	if strings.Trim(email, " ") == "" && password == "" {
		SignInPage(c)
		return
	}

	user, err := models.FindUserByEmail(email)

	if user != nil && user.Authenticate(password) {
		session.Set("user_id", user.ID)
		session.AddFlash("Log In Successfully!", "Info")
		session.Save()
		c.Redirect(http.StatusFound, "/")
	} else {
		if err != nil {
			log.Println("models.FindUserByEmail ==> ", err)
		}
		session.AddFlash("User name or Password Error.", "Warn")
		session.Save()
		c.Redirect(http.StatusFound, "/user")
	}
}

//LogOut 登出请求
func LogOut(c *gin.Context) {
	if isSignIn := c.GetBool("state_isUserSignIn"); !isSignIn {
		c.Redirect(http.StatusFound, "/")
		return
	}
	session := sessions.Default(c)
	session.Clear()
	session.AddFlash("Log Out Successfully!", "Info")
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

//Index 展示用户信息
func Index(c *gin.Context) {
	c.Redirect(http.StatusFound, "/user/sign_in")
}

//TODO：注册
