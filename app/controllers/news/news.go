package news

import (
	"gails/app/helpers"
	"gails/app/pkg/service"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

//Index 展示抓取消息
func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page == 0 {
		page = 1
	}

	hackerNews := service.HackerNewsService{
		Ctx:      c,
		URL:      "https://hacker-news.firebaseio.com/v0/",
		PageSize: 20,
	}

	newsID, err := hackerNews.GetTopStories(page)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "news", helpers.CommonHTMLRes("News", c))
		return
	}

	// 并发获取news
	var newsList []service.NewsItem
	var wg sync.WaitGroup
	ch := make(chan service.NewsItem)
	for _, id := range newsID {
		wg.Add(1)
		go hackerNews.GetItemGo(id, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for news := range ch {
		newsList = append(newsList, news)
	}

	sort.Slice(newsList, func(a, b int) bool {
		if newsList[a].Time < newsList[b].Time {
			return true
		}
		return false
	})

	data := gin.H{
		"newsList": newsList,
	}

	c.HTML(http.StatusOK, "news", helpers.CommonHTMLRes("News", c, data))
}
