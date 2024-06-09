package models

import "github.com/jinzhu/gorm"

type Certification struct {
	gorm.Model
	TutorID      uint
	Degree       string `gorm:"type:varchar(255)"`
	Experience   string `gorm:"type:text"`
	Certificates string `gorm:"type:json"`
	Tutor        Tutor
}
