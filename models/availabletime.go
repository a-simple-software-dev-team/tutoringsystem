package models

import "github.com/jinzhu/gorm"

type AvailableTime struct {
	gorm.Model
	TutorID   uint   `gorm:"not null"`
	Day       string `gorm:"type:varchar(10)"`
	TimeRange string `gorm:"type:varchar(50)"`
}
