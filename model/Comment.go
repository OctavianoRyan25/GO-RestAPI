package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID 	 uint 
	PhotoID  uint
	Message  string `gorm:"not null" json:"message" form:"message" valid:"required~Your message is required"`
	User	 *User
	Photo	 *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}