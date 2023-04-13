package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required~SocialMediaUrl is required"`
	UserID         uint   `json:"user_id" gorm:"not null"`
	User           *User  `gorm:"foreignKey:UserID"`
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
