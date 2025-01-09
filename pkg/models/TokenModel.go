package models

import (
	"MYSHOP/pkg/config"
	"gorm.io/gorm"
	"time"
)

type TokenModel struct {
	gorm.Model
	UserId int64     `json:"UserId"`
	Token  string    `json:"Token"`
	Expire time.Time `json:"Expire"`
	Ip     string    `json:"Ip"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&UserModel{})
	if err != nil {
		panic(err)
	}
}

func (t *TokenModel) CreateAccessToken() *TokenModel {
	db.Create(&t)
	return t
}
