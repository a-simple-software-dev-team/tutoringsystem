package models

import "github.com/jinzhu/gorm"

type Schedule struct {
	gorm.Model
	TutorID   uint   `gorm:"not null"`
	StudentID uint   `gorm:"not null"`
	Date      string `gorm:"type:date;not null"`
	Time      string `gorm:"type:time;not null"`
	Tutor     Tutor
	Student   Student
}
