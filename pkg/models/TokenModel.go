package models

import (
	"MYSHOP/pkg/config"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	UserId int64     `json:"UserId"`
	Token  string    `json:"Token"`
	Expire time.Time `json:"Expire"`
	Ip     string    `json:"Ip"`
	Device string    `json:"Device"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Token{})
	if err != nil {
		panic(err)
	}
}

func (t *Token) CreateAccessToken() *Token {
	db.Create(&t)
	return t
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Token{})
	if err != nil {
		panic(err)
	}
}

func TokenExists(token string) (uint, bool) {
	var result struct {
		UserID uint
	}

	err := db.Table("tokens").Where("token = ?", token).Select("user_id").Scan(&result).Error

	if err != nil {
		return 0, false
	}

	if result.UserID == 0 {
		return 0, false
	}

	return result.UserID, true
}
