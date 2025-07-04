package models

import (
	"gorm.io/gorm"
	"myshop/config/database"
)

type Category struct {
	gorm.Model
	NameUz         string `json:"NameUz" gorm:"type:varchar(255)"`
	NameRu         string `json:"NameRu" gorm:"type:varchar(255)"`
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

func CategoryCreate(c *Category) *Category {
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

func CategoryGetAllNoParents() []Category {
	var categories []Category
	db.Find(&categories)
	return categories
}

func CategoryGetById(id int64) (*Category, *gorm.DB) {
	var category Category
	db.Where("ID = ?", id).Find(&category)
	return &category, db
}

func CategoryUpdate(c Category) (int64, error) {
	re := db.Model(&c).Where("ID=?", c.ID).Updates(&c)
	return re.RowsAffected, re.Error
}

func CategoryDelete(id int64) Category {
	var category Category
	db.Where("ID = ?", id).Delete(&category)
	return category
}
