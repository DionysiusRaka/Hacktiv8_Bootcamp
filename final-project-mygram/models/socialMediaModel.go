package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Social media name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~SocialMediaUrl is required"`
	UserId         uint   `gorm:"user_id"`
	User           *User
}

func (sc *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sc)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (sc *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sc)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
