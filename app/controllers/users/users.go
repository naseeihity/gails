package users

import (
	"gails/app/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//SignInPage 登陆页
func SignInPage(c *gin.Context) {
	var isSignIn bool
	if isSignIn = c.GetBool("state_isUserSignIn"); isSignIn {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.HTML(http.StatusOK, "signin", gin.H{
		"Title":    "SignIn",
		"Active":   "signin",
		"IsSginIn": isSignIn,
	})
}

//SignIn 登陆请求
func SignIn(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if strings.Trim(email, " ") == "" && password == "" {
		SignInPage(c)
		return
	}

	user, err := models.FindUserByEmail(email)
	if err != nil {
		log.Println("models.FindUserByEmail ==> ", err)
		// TODO: FlashMessage
		SignInPage(c)
		return
	}
	if user != nil && user.Authenticate(password) {
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Save()
		c.Redirect(http.StatusFound, "/")
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
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

//Index 展示用户信息
func Index(c *gin.Context) {

}

//TODO：注册
