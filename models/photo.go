package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Title of your photo is required"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~PhotoUrl of your photo is required"`
	Caption  string `json:"caption" form:"caption" valid:"required~Caption of your product is required"`
	UserID   uint
	User     *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
