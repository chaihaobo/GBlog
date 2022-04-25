package category

import (
	"github.com/astaxie/beego/orm"
	"time"
)
import "api-core/models"

func Insert(category *models.Category) (bool, error) {
	o := orm.NewOrm()
	i, err := o.Insert(category)
	return i > 0, err
}

func SelectById(id int) *models.Category {
	entity := models.Category{
		BaseModel: models.BaseModel{
			Id: id,
		},
	}
	o := orm.NewOrm()
	err := o.Read(&entity)
	if err != nil {
		return nil
	}
	return &entity
}

func Save(category *models.Category) bool {
	o := orm.NewOrm()
	id := category.Id
	if id > 0 {
		entity := SelectById(id)
		if entity != nil {
			entity.Name = category.Name
			entity.UpdateTime = time.Now()
			_, err := o.Update(entity)
			if err != nil {
				return false
			}
		}
	} else {
		entity := models.Category{
			BaseModel: models.BaseModel{
				Id:         id,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			},
			Name:    category.Name,
			Article: nil,
		}
		_, err := o.Insert(&entity)
		if err != nil {
			return false
		}
	}
	return true
}

func List() []*models.Category {
	var categorys []*models.Category
	o := orm.NewOrm()
	o.QueryTable(new(models.Category)).All(&categorys)
	return categorys
}

func Delete(id int) bool {
	o := orm.NewOrm()
	_, error := o.Delete(&models.Category{BaseModel: models.BaseModel{Id: id}})
	if error != nil {
		return false
	}

	return true
}
