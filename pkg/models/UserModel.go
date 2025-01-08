package models

import (
	"MYSHOP/pkg/config"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type UserModel struct {
	gorm.Model
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Birthday  time.Time `json:"Birthday"`
	Phone     string    `json:"Phone" gorm:"unique"`
	Email     string    `json:"Email" gorm:"unique"`
	Role      string    `json:"Role"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&UserModel{})
	if err != nil {
		panic(err)
	}
}

func (u *UserModel) CreateUser() *UserModel {
	db.Create(&u)
	return u
}

func (u *UserModel) GetUserByPhone(phone string) *UserModel {
	var user UserModel
	db.Where("phone = ?", phone).First(&user)
	return &user
}

func (u *UserModel) GetUserByEmail(email string) *UserModel {
	var user UserModel
	db.Where("email = ?", email).First(&user)
	return &user
}

func GetAllUsers() []UserModel {
	var users []UserModel
	db.Find(&users)
	return users
}

func GetUserById(Id int64) (*UserModel, *gorm.DB) {
	var getUser UserModel
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func UpdateUser(user UserModel) (int64, error) {
	re := db.Model(&user).Where("ID=?", user.ID).Updates(&user)
	return re.RowsAffected, re.Error
}
