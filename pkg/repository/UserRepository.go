package repository

type UserLoginStructRepository struct {
	Phone    string `json:"Phone" validate:"required,numeric,min=9,max=22"`
	Password string `json:"Password" validate:"required,min=4,max=22"`
}
