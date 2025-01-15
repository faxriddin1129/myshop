package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"myshop/config/database"
	"myshop/helpers"
	"time"
)

var db *gorm.DB

type User struct {
	gorm.Model
	FirstName    string    `json:"FirstName" gorm:"type:varchar(255)"`
	LastName     string    `json:"LastName" gorm:"type:varchar(255)"`
	Birthday     time.Time `json:"Birthday"`
	Phone        string    `json:"Phone" gorm:"unique" gorm:"type:varchar(22)"`
	Email        string    `json:"Email" gorm:"unique" gorm:"type:varchar(255)"`
	Role         int8      `json:"Role"`
	PasswordHash string    `json:"PasswordHash" gorm:"type:varchar(255)"`
	Status       int8      `json:"Status"`
	SessionId    string    `json:"SessionId" gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "users"
}

func init() {
	database.Connect()
	db = database.GetDB()
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	defaultAdmin := User{
		FirstName:    "Fakhriddin",
		LastName:     "Boboyev",
		Birthday:     time.Date(2000, time.April, 6, 0, 0, 0, 0, time.UTC),
		Phone:        "998907291129",
		Email:        "fakhriddin1129@gmail.com",
		Role:         helpers.ROLE_ADMIN,
		PasswordHash: HashPassword("fakhriddin1129"),
		Status:       1,
	}

	var existingAdmin User
	db.Where("email = ?", defaultAdmin.Email).First(&existingAdmin)
	if existingAdmin.ID == 0 {
		if err := db.Create(&defaultAdmin).Error; err != nil {
			panic("Admin Create Admin: " + err.Error())
		}
	}
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

func GetUserByPhone(Phone string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("Phone=?", Phone).Find(&getUser)
	return &getUser, db
}

func GetUserByEmail(Email string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("Email=?", Email).Find(&getUser)
	return &getUser, db
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func UpdateUser(user User) (int64, error) {
	re := db.Model(&user).Where("ID=?", user.ID).Updates(&user)
	return re.RowsAffected, re.Error
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
