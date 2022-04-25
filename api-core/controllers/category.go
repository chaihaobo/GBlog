package controllers

import (
	"api-core/models"
	"api-core/models/category"
)

type CategoryController struct {
	BaseController
}

// @router /save [post]
func (this *CategoryController) Save() {
	var category models.Category
	this.parseBody(&category)
	this.validation(&category)
	this.success()
}

// @router /list [get]
func (this *CategoryController) List() {
	this.successData(category.List())
}
