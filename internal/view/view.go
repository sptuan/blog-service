package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type View struct {
}

func NewView() View {
	return View{}
}

func (v View) ArticleList(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{})
}
