package models

import (
	"gorm.io/gorm"
	"myshop/config/database"
)

type ConfigModel struct {
	gorm.Model
	Phone        string `json:"Phone"`
	Email        string `json:"Email"`
	Telegram     string `json:"Telegram"`
	TelegramNick string `json:"TelegramNick"`
	Instagram    string `json:"Instagram"`
	Facebook     string `json:"Facebook"`
	Youtube      string `json:"Youtube"`
	Address      string `json:"Address"`
	PublicOffer  string `json:"PublicOffer"`
	FooterText   string `json:"FooterText"`
}

func (ConfigModel) TableName() string { return "config" }

func init() {
	database.Connect()
	db := database.GetDB()
	err := db.AutoMigrate(&ConfigModel{})
	if err != nil {
		panic(err)
	}
}
