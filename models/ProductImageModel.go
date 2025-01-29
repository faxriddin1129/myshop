package models

import (
	"gorm.io/gorm"
	"myshop/config/database"
)

type ProductImage struct {
	gorm.Model
	ProductId  int64  `json:"ProductId"`
	FileId     int64  `json:"FileId"`
	CurrentUrl string `json:"CurrentUrl"`
}

func (ProductImage) TableName() string {
	return "product_images"
}

func init() {
	database.Connect()
	db = database.GetDB()
	err := db.AutoMigrate(&ProductImage{})
	if err != nil {
		panic(err)
	}
}

func GetProductImages(productId int64) ([]ProductImage, error) {
	var productImages []ProductImage
	db.Where("product_id = ?", productId).Find(&productImages)
	return productImages, nil
}
