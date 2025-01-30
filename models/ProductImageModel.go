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

func GetProductImageById(Id int64) (*ProductImage, *gorm.DB) {
	var productImage ProductImage
	db := db.Where("ID=?", Id).Find(&productImage)
	return &productImage, db
}

func CreateProductImages(pi *ProductImage) *ProductImage {
	db.Create(pi)
	return pi
}

func ProductImagesDelete(id int64) ProductImage {
	var productImage ProductImage
	db.Where("ID = ?", id).Delete(&productImage)
	return productImage
}
