package models

import (
	"gorm.io/gorm"
	"myshop/config/database"
)

type Brand struct {
	gorm.Model
	NameUz         string `json:"NameUz" gorm:"type:varchar(255)"`
	NameRu         string `json:"NameRu" gorm:"type:varchar(255)"`
	FileId         int64  `json:"FileId"`
	CurrentFileUrl string `json:"CurrentFileUrl" gorm:"type:varchar(255)"`
}

func (Brand) TableName() string {
	return "brands"
}

func init() {
	database.Connect()
	db = database.GetDB()
	err := db.AutoMigrate(&Brand{})
	if err != nil {
		panic(err)
	}
}

func BrandCreate(c *Brand) *Brand {
	db.Create(&c)
	return c
}

func BrandGetAll() []Brand {
	var brands []Brand
	db.Find(&brands)
	return brands
}

func BrandGetById(id int64) (*Brand, *gorm.DB) {
	var brand Brand
	db.Where("ID = ?", id).Find(&brand)
	return &brand, db
}

func BrandUpdate(c Brand) (int64, error) {
	res := db.Model(&c).Where("ID=?", c.ID).Updates(&c)
	return res.RowsAffected, res.Error
}

func BrandDelete(id int64) Brand {
	var brand Brand
	db.Where("ID = ?", id).Delete(&brand)
	return brand
}
