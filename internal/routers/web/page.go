package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sptuan/blog-service/internal/view"
)

type Page struct {
}

func NewPage() Page {
	return Page{}
}

func (p Page) Index(c *gin.Context) {
	// 1. Create a new view
	v := view.NewView()
	// 2. Data Modify/Query here (DAO)

	// 3. Response HTML
	v.ArticleList(c)
}
