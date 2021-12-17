package dto

import "time"

type MarkerDTO struct {
	Id 	   		uint	 	`json:"id" binding:"-"`
	Longitude string 		`json:"longitude" binding:"required"`
	Latitude  string 		`json:"latitude"  binding:"required"`
	Name      string 		`json:"name" binding:"required"`
	Address   string 		`json:"address" binding:"-"`
	CreatedAt time.Time 	`json:"createdAt" binding:"-"`
	UpdatedAt time.Time 	`json:"updatedAt" binding:"-"`
}
