package models

import "github.com/jinzhu/gorm"

type Notification struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Message string `gorm:"type:text"`
	Read    bool   `gorm:"default:false"`
	User    User
}
