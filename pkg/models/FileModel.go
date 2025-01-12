package models

import (
	"MYSHOP/pkg/config"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	BaseUrl    string `json:"BaseUrl" gorm:"type:varchar(255)"`
	Path       string `json:"Path" gorm:"type:varchar(255)"`
	CurrentUrl string `json:"CurrentUrl" gorm:"type:varchar(255)"`
	Name       string `json:"Name" gorm:"type:varchar(255)"`
}

func (File) TableName() string {
	return "files"
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&File{})
	if err != nil {
		panic(err)
	}
}
