package dto

type ArticleDTO struct {
	Id         int    `json:"id"`
	Title      string `json:"title" valid:"Required"`
	Content    string `json:"content" valid:"Required"`
	CategoryId int    `json:"categoryId"`
}
