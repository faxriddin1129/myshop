package models

import (
	"gorm.io/gorm"
	"myshop/config/database"
)

type Banner struct {
	gorm.Model
	ProductId            int64  `json:"ProductId"`
	FileId               int64  `json:"FileId"`
	MobileFileId         int64  `json:"FileIdMobile"`
	CurrentFileUrl       string `json:"CurrentFileUrl" gorm:"type:varchar(255)"`
	MobileCurrentFileUrl string `json:"MobileCurrentFileUrl" gorm:"type:varchar(255)"`
}

func (Banner) TableName() string {
	return "banners"
}

func init() {
	database.Connect()
	db = database.GetDB()
	err := db.AutoMigrate(&Banner{})
	if err != nil {
		panic(err)
	}
}

func BannerCreate(c *Banner) *Banner {
	db.Create(&c)
	return c
}

func BannerGetAll() []Banner {
	var banners []Banner
	db.Find(&banners)
	return banners
}

func BannerGetById(id int64) (*Banner, *gorm.DB) {
	var banner Banner
	db.Where("ID = ?", id).Find(&banner)
	return &banner, db
}

func BannerDelete(id int64) Banner {
	var banner Banner
	db.Where("ID = ?", id).Delete(&banner)
	return banner
}
