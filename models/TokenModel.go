package models

import (
	"gorm.io/gorm"
	"myshop/config/database"
	"time"
)

type Token struct {
	gorm.Model
	UserId int64     `json:"UserId"`
	Token  string    `json:"Token" gorm:"type:varchar(255)"`
	Expire time.Time `json:"Expire"`
	Ip     string    `json:"Ip" gorm:"type:varchar(255)"`
	Device string    `json:"Device"`
}

func (Token) TableName() string {
	return "tokens"
}

func init() {
	database.Connect()
	db = database.GetDB()
	err := db.AutoMigrate(&Token{})
	if err != nil {
		panic(err)
	}
}

func (t *Token) CreateAccessToken() *Token {
	db.Create(&t)
	return t
}

func TokenExists(token string) (uint, bool) {
	var result struct {
		UserID uint
	}

	err := db.Table("tokens").Where("token = ? AND expire > ?", token, time.Now()).Select("user_id").Scan(&result).Error

	if err != nil {
		return 0, false
	}

	if result.UserID == 0 {
		return 0, false
	}

	return result.UserID, true
}
