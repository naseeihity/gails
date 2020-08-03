package service

import (
	"encoding/json"
	"fmt"
	"gails/app/helpers"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

//HackerNewsService :
type HackerNewsService struct {
	Ctx      *gin.Context
	URL      string
	PageSize int
}

//NewsOptions :
type NewsOptions struct {
	CacheKey    string
	CacheExpire int
}

//NewsItem :
type NewsItem struct {
	ID          int    `json:"id"`
	Deleted     bool   `json:"deleted"`
	Type        string `json:"type"`
	By          string `json:"by"`
	Time        int    `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Parent      int    `json:"parent"`
	Kids        []int  `json:"kids"`
	URL         string `json:"url"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Parts       []int  `json:"parts"`
	Descendants int    `json:"descendants"`
}

//Request :请求封装，设置和获取缓存
func (h *HackerNewsService) Request(api string, options NewsOptions) ([]byte, error) {
	if options.CacheKey != "" {
		// get from cache
	}

	url := fmt.Sprintf("%s/%s", h.URL, api)
	resp, err := helpers.GetRawRes(url, nil)
	if err != nil {
		log.Println("Get news Failed:", err)
		return nil, err
	}

	if options.CacheKey != "" && len(resp) > 0 {
		// set cache
	}

	return resp, nil
}

//GetTopStories : 获取最新stroies
func (h *HackerNewsService) GetTopStories(p int) ([]int, error) {
	var stories []int
	const PAGESIZE = 20
	page := p
	if page <= 1 {
		page = 1
	}
	opt := NewsOptions{
		CacheKey:    fmt.Sprintf("new_ids_%d", page),
		CacheExpire: 300,
	}
	res, err := h.Request("topstories.json", opt)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(res, &stories)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var startAt int
	var endAt int
	if page > 1 {
		startAt = (page-1)*PAGESIZE - 1
	}
	endAt = page * PAGESIZE
	return stories[startAt:endAt], nil
}

//GetItem :获取详情
func (h *HackerNewsService) GetItem(id int) (NewsItem, error) {
	var news NewsItem

	api := fmt.Sprintf("item/%d.json", id)
	opt := NewsOptions{
		CacheKey:    fmt.Sprintf("new_item_%d", id),
		CacheExpire: 300,
	}
	res, err := h.Request(api, opt)
	if err != nil {
		log.Println(err)
		return news, err
	}

	err = json.Unmarshal(res, &news)
	return news, err
}

//GetItemGo 并发获取详情
func (h *HackerNewsService) GetItemGo(id int, ch chan<- NewsItem, wg *sync.WaitGroup) {
	defer wg.Done()
	news, err := h.GetItem(id)
	if err != nil {
		log.Println(err)
	}
	ch <- news
}
