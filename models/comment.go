package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint   `json:"user_id" gorm:"not null"`
	PhotoID uint   `json:"photo_id" gorm:"not null"`
	Message string `json:"message" form:"message" valid:"required~Message is required"`
	User    *User  `gorm:"foreignKey:UserID"`
	Photo   *Photo `gorm:"foreignKey:PhotoID"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
