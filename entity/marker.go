package entity

import (
	"time"
)

type Marker struct {
	Id 	   		uint   `gorm:"primaryKey"`
	Longitude string
	Latitude  string
	Name    	string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (marker *Marker) TableName() string {
	return "markers"
}
