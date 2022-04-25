package controllers

import (
	"api-core/models"
	categoryModel "api-core/models/category"
	"strconv"
)

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Save() {
	var category models.Category
	c.parseBody(&category)
	c.validation(&category)
	categoryModel.Save(&category)
	c.success()
}

func (c *CategoryController) List() {
	c.successData(categoryModel.List())
}

func (c *CategoryController) Delete() {
	param := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(param)
	categoryModel.Delete(id)
	c.success()
}
