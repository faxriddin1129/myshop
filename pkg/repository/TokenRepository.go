package repository

type TokenRepository struct {
	UserId int64  `json:"UserId" validate:"required"`
	Token  string `json:"Token" validate:"required"`
	Expire string `json:"Expire" validate:"required"`
	Ip     string `json:"Ip" validate:"required"`
}
