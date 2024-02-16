package models

import (
	"Hactiv8_Bootcamp/final-project-mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email    string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required, email~Invalid email format"`
	Password string        `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)~Password has to have minimum length of 6 characters"`
	Age      int           `gorm:"not null" json:"age" form:"age" valid:"required~Age is required"`
	Photos   []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comments []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
