package models

import (
	"MYSHOP/config"
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

func CreateFileModel(f *File) *File {
	db.Create(f)
	return f
}

func GetFileById(Id int64) (*File, *gorm.DB) {
	var getFile File
	db := db.Where("ID=?", Id).Find(&getFile)
	return &getFile, db
}
