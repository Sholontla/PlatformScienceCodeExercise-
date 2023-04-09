package entity

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
}

func (u *User) Validations(user string) error {
	if u.Username == user {
		err := fmt.Errorf("missing user parameter")
		return err
	}
	return nil
}
