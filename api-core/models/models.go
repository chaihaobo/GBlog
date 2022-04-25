package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	// register model
	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(Article))
}

type BaseModel struct {
	Id         int       `orm:"auto" json:"id"`
	CreateTime time.Time `gorm:"type:time" json:"createTime,omitempty"`
	UpdateTime time.Time `gorm:"type:time" json:"updateTime,omitempty"`
}

type Category struct {
	BaseModel
	Name    string     `json:"name" valid:"Required"`
	Article []*Article `orm:"reverse(many)" `
}

type Article struct {
	BaseModel
	Title    string    `json:"title" valid:"Required"`
	Content  string    `json:"content" valid:"Required"`
	Category *Category `orm:"rel(fk)"`
}
