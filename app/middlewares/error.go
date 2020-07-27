package middlewares

import (
	"github.com/gin-gonic/gin"
)

//ErrorHandle 全局统一的错误码处理
func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Next()
		// err := c.Errors.Last()
		// if err == nil {
		// 	return
		// }

		// // Use reflect.TypeOf(err.Err) to known the type of your error
		// if error, ok := errors.Cause(err.Err).(*myspace.KindOfClientError); ok {
		// 	c.JSON(400, gin.H{
		// 		"error": "Blah blahhh",
		// 	})
		// 	return
		// }
	}
}
