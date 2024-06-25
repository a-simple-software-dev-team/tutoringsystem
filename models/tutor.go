package models

import "github.com/jinzhu/gorm"

type Tutor struct {
	gorm.Model
	UserID         uint   `gorm:"not null"`
	Subjects       string `gorm:"type:varchar(255)"`
	Grades         string `gorm:"type:varchar(255)"`
	User           User
	AvailableTimes []AvailableTime `gorm:"foreignkey:TutorID"`
}
