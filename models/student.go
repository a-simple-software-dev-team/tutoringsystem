package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	UserID         uint   `gorm:"not null"`
	SubjectsNeeded string `gorm:"type:varchar(255)"`
	Grades         string `gorm:"type:varchar(255)"`
	User           User
	AvailableTimes []AvailableTime `gorm:"foreignkey:StudentID"`
}
