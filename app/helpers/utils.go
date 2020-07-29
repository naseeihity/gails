package helpers

import (
	"gails/app/models"

	"github.com/gin-gonic/gin"
)

func isActive(action string, param string) string {
	var cls string
	if action == param {
		cls = "active"
	}
	return cls
}

//InterArrToStrArr interface 数组转为字符串数组
func InterArrToStrArr(interArr []interface{}) []string {
	var strArr []string
	for _, inter := range interArr {
		strArr = append(strArr, inter.(string))
	}
	return strArr
}

//CommonHTMLRes common with params
func CommonHTMLRes(active string, c *gin.Context, params ...gin.H) gin.H {
	var username string
	isSignIn := c.GetBool("state_isUserSignIn")
	if user, _ := c.Get("state_curUser"); user != nil {
		if curUser, ok := user.(*models.Users); ok {
			username = curUser.Name
		}
	}

	common := gin.H{
		"Title":    active,
		"Active":   active,
		"IsSginIn": isSignIn,
		"curUser":  username,
		"Msg":      c.MustGet("flashMessage"),
	}

	params = append(params, common)

	res := MergeMaps(params...)

	return res
}

//MergeMaps 合并gin.H
func MergeMaps(maps ...gin.H) gin.H {
	result := make(gin.H)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
