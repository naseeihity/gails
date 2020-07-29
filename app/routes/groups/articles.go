package groups

import (
	"gails/app/controllers/articles"

	"github.com/gin-gonic/gin"
)

//InitArticles :
func InitArticles(r *gin.Engine) {
	a := r.Group("/articles")

	// for RESTful API support problem:https://github.com/gin-gonic/gin/issues/2016
	// The routers is not written in RESTful style
	{
		a.GET("/new", articles.NewArticle)
		a.GET("/detail/:id", articles.Show)
		// a.Use(middlewares.CheckLogin())
		a.PUT("/update/:id", articles.CheckArticleOwner(), articles.CheckParamsBody(), articles.Update)
		a.GET("/edit/:id", articles.CheckArticleOwner(), articles.Edit)
		a.POST("/", articles.CheckParamsBody(), articles.Create)

	}

}
