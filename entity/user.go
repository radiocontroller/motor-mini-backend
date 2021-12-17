package entity

import (
	"time"
)

type User struct {
	Id 	   		uint   `gorm:"primaryKey"`
	Openid 	  string
	Unionid   string
	SessionKey string
	Token   	string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) TableName() string {
	return "users"
}
