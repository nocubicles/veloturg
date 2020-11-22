package models

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"index"`
	Ads         []Ad
	Sessions    []Session
	SignInToken uuid.UUID `gorm:"index"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.SignInToken, err = uuid.NewV4()

	if err != nil {
		return err
	}

	return
}
