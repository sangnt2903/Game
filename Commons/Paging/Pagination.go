package Paging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
	"strings"
)

// Paginator
type Paginator struct {
	RootPath    string      `json:"-"`
	TotalRecord int64       `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	Sort        string      `json:"sort"`
	PrevPage    string      `json:"prev_page"`
	NextPage    string      `json:"next_page"`
}

func NewPaginator(c *gin.Context) *Paginator {
	var limit int = 25
	var page int = 1
	sort := "uuid ASC"

	query := c.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}

	return &Paginator{
		RootPath: c.Request.URL.Path,
		Page:     page,
		Limit:    limit,
		Sort:     sort,
	}
}

// Paging
func (p *Paginator) Paging(db *gorm.DB, result interface{}) error {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if p.Sort != "" {
		sortItems := strings.Split(p.Sort, "|")
		for _, o := range sortItems {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var count int64
	var offset int

	countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	if err := db.Limit(p.Limit).Offset(offset).Find(result).Error; err != nil {
		return err
	}

	p.TotalRecord = count
	p.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		p.PrevPage = fmt.Sprintf("%s?page=%d&limit=%d&sort=%s", p.RootPath, p.Page-1, p.Limit, p.Sort)

	} else {
		p.PrevPage = fmt.Sprintf("%s?page=%d&limit=%d&sort=%s", p.RootPath, p.Page, p.Limit, p.Sort)
	}

	if p.Page == p.TotalPage {
		p.NextPage = fmt.Sprintf("%s?page=%d&limit=%d&sort=%s", p.RootPath, p.Page, p.Limit, p.Sort)
	} else {
		p.NextPage = fmt.Sprintf("%s?page=%d&limit=%d&sort=%s", p.RootPath, p.Page+1, p.Limit, p.Sort)
	}
	return nil
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int64) {
	db.Model(anyType).Count(count)
	done <- true
}
