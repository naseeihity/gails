package articles

import "github.com/gin-gonic/gin"

//NewArticle 新建文章页面
func NewArticle(c *gin.Context) {

}

//Show 文章详情
func Show(c *gin.Context) {

}

//Update 更新文章
func Update(c *gin.Context) {

}

//Edit 编辑页面
func Edit(c *gin.Context) {

}

//Create 新建文件
func Create(c *gin.Context) {

}

// ====== Middlewares ======

//CheckArticleOwner 检查文章权限
func CheckArticleOwner() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//CheckParamsBody 检查参数
func CheckParamsBody() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
