package groups

import (
	"gails/app/controllers/articles"
	"gails/app/middlewares"

	"github.com/gin-gonic/gin"
)

//InitArticles :
func InitArticles(r *gin.Engine) {
	a := r.Group("/articles")

	{
		a.GET("/new", articles.NewArticle)
		a.GET("/:id", articles.Show)
		a.Use(middlewares.CheckLogin())
		a.PUT("/:id", articles.CheckArticleOwner(), articles.CheckParamsBody(), articles.Update)
		a.GET("/:id/edit", articles.CheckArticleOwner(), articles.Edit)
		a.POST("/", articles.CheckParamsBody(), articles.Create)
	}

}
