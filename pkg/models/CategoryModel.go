package models

import (
	"MYSHOP/config/database"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	NameUz         string `json:"NameUz" gorm:"type:varchar(255)"`
	NameRu         string `json:"NameRu" gorm:"type:varchar(255)"`
	ImageUrl       string `json:"ImageUrl" gorm:"type:varchar(255)"`
	ParentID       *int64 `json:"ParentID"`
	Type           string `json:"Type" gorm:"type:varchar(255)"`
	FileId         int64  `json:"FileId"`
	CurrentFileUrl string `json:"CurrentFileUrl" gorm:"type:varchar(255)"`
}

func (Category) TableName() string {
	return "categories"
}

func init() {
	database.Connect()
	db = database.GetDB()
	err := db.AutoMigrate(&Category{})
	if err != nil {
		panic(err)
	}
}

func CreateCategory(c *Category) *Category {

	if c.ParentID != nil && *c.ParentID == 0 {
		c.ParentID = nil
	}
	db.Create(&c)
	return c
}

func CategoryGetAll(parent int) []Category {
	var categories []Category
	if parent > 0 {
		db.Where("parent_id = ?", parent).Find(&categories)
	} else {
		db.Where("parent_id IS NULL").Find(&categories)
	}
	return categories
}
