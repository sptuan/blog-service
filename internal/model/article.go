package model

import "github.com/sptuan/blog-service/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desciption    string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article_tag"
}
