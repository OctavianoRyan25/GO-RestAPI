package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title	 string `gorm:"not null" json:"title" form:"title" valid:"required~Your title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Your photo_url is required"`
	UserID	 uint
	User	 *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}