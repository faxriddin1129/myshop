package models

import (
	"gorm.io/gorm"
	"myshop/config/database"
	"time"
)

type Product struct {
	gorm.Model
	NameUz             string    `json:"NameUz" gorm:"type:varchar(255)"`
	NameRu             string    `json:"NameRu" gorm:"type:varchar(255)"`
	FileId             int64     `json:"FileId"`
	CurrentFileUrl     string    `json:"CurrentFileUrl" gorm:"type:varchar(255)"`
	ShortDescriptionUz string    `json:"ShortDescriptionUz"`
	ShortDescriptionRu string    `json:"ShortDescriptionRu"`
	DescriptionUz      string    `json:"DescriptionUz"`
	DescriptionRu      string    `json:"DescriptionRu"`
	Count              int       `json:"Count"`
	Price              float64   `json:"Price"`
	DiscountPrice      int       `json:"DiscountPrice"`
	Status             int       `json:"Status"`
	CategoryId         int       `json:"CategoryId"`
	BrandId            int       `json:"BrandId"`
	ProductionTime     time.Time `json:"ProductionTime"`
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

func (p *Product) ProductCreate() *Product {
	db.Create(&p)
	return p
}

func ProductDelete(id int64) Product {
	var product Product
	db.Where("ID = ?", id).Delete(&product)
	return product
}
