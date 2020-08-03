package helpers

import (
	"gails/app/models"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

//GetRawRes 获取raw响应
func GetRawRes(url string, params map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 设置Get请求的params
	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// 响应体转为string
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return bodyBytes, nil
	}

	return nil, err
}
