package models

import (
	"gorm.io/gorm"
	"myshop/config/database"
)

type Config struct {
	gorm.Model
	Phone        string `json:"Phone" gorm:"type:varchar(255); default:null"`
	Email        string `json:"Email" gorm:"type:varchar(255); default:null"`
	Telegram     string `json:"Telegram" gorm:"type:varchar(255); default:null"`
	TelegramNick string `json:"TelegramNick" gorm:"type:varchar(255); default:null"`
	Instagram    string `json:"Instagram" gorm:"type:varchar(255); default:null"`
	Facebook     string `json:"Facebook" gorm:"type:varchar(255); default:null"`
	Youtube      string `json:"Youtube" gorm:"type:varchar(255); default:null"`
	Address      string `json:"Address" gorm:"default:null"`
	PublicOffer  string `json:"PublicOffer" gorm:"type:varchar(255); default:null"`
	FooterText   string `json:"FooterText" gorm:"default:null"`
}

type ConfigValidate struct {
	Phone        string `json:"Phone" validate:"required"`
	Email        string `json:"Email" validate:"required"`
	Telegram     string `json:"Telegram" validate:"max=255"`
	TelegramNick string `json:"TelegramNick" validate:"max=255"`
	Instagram    string `json:"Instagram" validate:"max=255"`
	Facebook     string `json:"Facebook" validate:"max=255"`
	Youtube      string `json:"Youtube" validate:"max=255"`
	Address      string `json:"Address" validate:"max=1000"`
	PublicOffer  string `json:"PublicOffer" validate:"max=255"`
	FooterText   string `json:"FooterText" validate:"max=1000"`
}

func (Config) TableName() string { return "config" }

func init() {
	database.Connect()
	db := database.GetDB()
	err := db.AutoMigrate(&Config{})
	if err != nil {
		panic(err)
	}

	defaultModel := Config{
		Phone: "+998907291129",
		Email: "fakhriddin1129@gmail.com",
	}
	var existingConfig Config
	db.Where("ID = ?", 1).First(&existingConfig)
	if existingConfig.ID == 0 {
		if err := db.Create(&defaultModel).Error; err != nil {
			panic("Admin Create Admin: " + err.Error())
		}
	}
}

func ConfigGetById(Id int64) (*Config, *gorm.DB) {
	var config Config
	db := db.Where("ID=?", Id).Find(&config)
	return &config, db
}

func ConfigUpdate(c Config) (int64, error) {
	res := db.Model(&c).Where("ID=?", c.ID).Updates(c)
	return res.RowsAffected, res.Error
}
