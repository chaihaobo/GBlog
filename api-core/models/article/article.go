package article

import (
	"api-core/models"
	"api-core/models/category"
	"api-core/models/dto"
	"github.com/astaxie/beego/orm"
	"time"
)

func SelectById(id int) *models.Article {
	entity := models.Article{
		BaseModel: models.BaseModel{
			Id: id,
		},
	}
	o := orm.NewOrm()
	err := o.Read(&entity)
	if err != nil {
		return nil
	}
	o.LoadRelated(&entity, "Category")
	return &entity
}

func ListByCategoryId(categoryId int) []*models.Article {
	var articleList []*models.Article
	o := orm.NewOrm()
	if categoryId > 0 {
		o.QueryTable(new(models.Article)).Filter("category__id", categoryId).OrderBy("createTime").All(&articleList)
	} else {
		o.QueryTable(new(models.Article)).OrderBy("createTime").All(&articleList)
	}

	return articleList
}

func Save(dto *dto.ArticleDTO) bool {
	o := orm.NewOrm()
	id := dto.Id
	cat := category.SelectById(dto.CategoryId)
	if id > 0 {
		entity := SelectById(id)
		if entity != nil {
			entity.UpdateTime = time.Now()
			entity.Title = dto.Title
			entity.Content = dto.Content
			entity.Category = cat
			if _, err := o.Update(entity); err != nil {
				return false
			}
		}
	} else {
		entity := models.Article{
			BaseModel: models.BaseModel{
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			},
			Title:    dto.Title,
			Content:  dto.Content,
			Category: cat,
		}
		if _, err := o.Insert(&entity); err != nil {
			return false
		}
	}
	return true

}
