package models

import "github.com/jinzhu/gorm"

//Articles 文章结构体
type Articles struct {
	gorm.Model
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	Content  string `json:"content"`
	UserID   int    `json:"userID"`
}

//ArticleInfo article with username
type ArticleInfo struct {
	Articles
	Name string
}

//GetArticlesByPage 分页查询
func GetArticlesByPage(page int) (*[]ArticleInfo, error) {
	var articles []ArticleInfo
	pageOffset := (page - 1) * 10
	res := db.Table("articles").Select("articles.title, articles.created_at, articles.id, articles.abstract, users.name").Joins("left join users on users.id = articles.user_id").Limit(10).Offset(pageOffset).Order("created_at desc").Scan(&articles)

	return &articles, res.Error
}

//GetArticlesCount 查询文章总数
func GetArticlesCount() (int, error) {
	var count int
	res := db.Model(&Articles{}).Count(&count)

	return count, res.Error
}

//GetArticleByID 根据id查询文章
func GetArticleByID(ID int) (*Articles, error) {
	var article Articles
	res := db.First(&article, ID)

	return &article, res.Error
}
