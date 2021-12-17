package dto

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserDTO struct {
	Openid 	  string	`json:"openid"`
	Unionid   string  `json:"unionid"`
	SessionKey string `json:"session_key"`
	Token   	string  `json:"token"`
}

func (u *UserDTO) BeforeCreate(tx *gorm.DB) (err error) {
	u.Token = uuid.NewV4().String()
	return
}
