package repository

import (
	"net/http"
	"time"
)

type ProductValidateRepository struct {
	NameUz             string    `json:"NameUz" validate:"required,max=255"`
	NameRu             string    `json:"NameRu" validate:"required,max=255"`
	FileId             int64     `json:"FileId" validate:"required"`
	CurrentFileUrl     string    `json:"CurrentFileUrl" validate:"omitempty,max=255"`
	ShortDescriptionUz string    `json:"ShortDescriptionUz" validate:"required,max=500"`
	ShortDescriptionRu string    `json:"ShortDescriptionRu" validate:"required,max=500"`
	DescriptionUz      string    `json:"DescriptionUz" validate:"omitempty,max=10000"`
	DescriptionRu      string    `json:"DescriptionRu" validate:"omitempty,max=10000"`
	Count              int       `json:"Count" validate:"required,gte=0"`
	Price              float64   `json:"Price" validate:"required,gte=0"`
	DiscountPrice      int       `json:"DiscountPrice" validate:"gte=0"`
	Status             int       `json:"Status" validate:"required,oneof=0 1"`
	CategoryId         int       `json:"CategoryId" validate:"required"`
	BrandId            int       `json:"BrandId" validate:"required"`
	ProductionTime     time.Time `json:"ProductionTime" validate:"omitempty"`
}

func ProductSave(w http.ResponseWriter, r *http.Request) {

}
