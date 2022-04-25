package controllers

import "api-core/models/dto"
import "api-core/models/article"

type ArticleController struct {
	BaseController
}

func (c *ArticleController) save() {
	articleDTO := dto.ArticleDTO{}
	c.parseBody(&articleDTO)
	c.validation(&articleDTO)
	article.Save(&articleDTO)
	c.success()
}
