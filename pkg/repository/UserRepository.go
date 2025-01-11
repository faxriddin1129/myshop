package repository

import "time"

type UserLoginStructRepository struct {
	Phone    string `json:"Phone" validate:"required,numeric,min=9,max=12"`
	Password string `json:"Password" validate:"required,min=4,max=22"`
}

type UserSendSmsStructRepository struct {
	Phone  string    `json:"Phone" validate:"required,numeric,min=9,max=12"`
	Token  string    `json:"Token" validate:"required,string"`
	Expire time.Time `json:"Expire"`
}
