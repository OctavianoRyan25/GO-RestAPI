package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name	 		string `gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	SocialMediaUrl 	string `gorm:"not null" json:"social_media_url" form:"socual_media_url" valid:"required~Your social_media_url is required"`
	UserID	 		uint
	User	 		*User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}