package controllers

import (
	"api-core/models/dto"
	"strconv"
)
import "api-core/models/article"

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Save() {
	articleDTO := dto.ArticleDTO{}
	c.parseBody(&articleDTO)
	c.validation(&articleDTO)
	article.Save(&articleDTO)
	c.success()
}

func (c *ArticleController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	result := article.SelectById(id)
	c.successData(result)
}

func (c *ArticleController) List() {
	categoryId, _ := c.GetInt("categoryId", 0)
	c.successData(article.ListByCategoryId(categoryId))

}
