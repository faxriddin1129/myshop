package models

import (
	"gorm.io/gorm"
	"myshop/config/database"
)

type Product struct {
	gorm.Model
	NameUz             string  `json:"NameUz" gorm:"type:varchar(255)"`
	NameRu             string  `json:"NameRu" gorm:"type:varchar(255)"`
	FileId             int64   `json:"FileId"`
	CurrentFileUrl     string  `json:"CurrentFileUrl" gorm:"type:varchar(255)"`
	ShortDescriptionUz string  `json:"ShortDescriptionUz"`
	ShortDescriptionRu string  `json:"ShortDescriptionRu"`
	DescriptionUz      string  `json:"DescriptionUz"`
	DescriptionRu      string  `json:"DescriptionRu"`
	Count              int     `json:"Count"`
	Price              float64 `json:"Price"`
	DiscountPrice      float64 `json:"DiscountPrice"`
	Status             int8    `json:"Status"`
	CategoryId         int64   `json:"CategoryId"`
	BrandId            int64   `json:"BrandId"`
	ProductionTime     string  `json:"ProductionTime" gorm:"type:varchar(255)"`
}

func (Product) TableName() string {
	return "products"
}

func init() {
	database.Connect()
	db = database.GetDB()
	err := db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
}

func ProductCreate(p *Product) *Product {
	db.Create(&p)
	return p
}

func ProductGetAll() []Product {
	var products []Product
	db.Find(&products)
	return products
}

func ProductGetById(id int64) (*Product, *gorm.DB) {
	var product Product
	db.Where("ID = ?", id).Find(&product)
	return &product, db
}

func ProductUpdate(c Product) (int64, error) {
	re := db.Model(&c).Where("ID=?", c.ID).Updates(&c)
	return re.RowsAffected, re.Error
}

func ProductDelete(id int64) Product {
	var product Product
	db.Where("ID = ?", id).Delete(&product)
	return product
}
