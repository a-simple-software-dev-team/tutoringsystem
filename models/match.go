package models

import "github.com/jinzhu/gorm"

type Match struct {
	gorm.Model
	UserID             uint    `gorm:"not null"`
	MatchID            uint    `gorm:"not null"`
	CompatibilityScore float64 `gorm:"type:float"`
	Role               string  `gorm:"not null"`
	User               User
}
