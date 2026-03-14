package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	UUID    string `gorm:"type:char(36);uniqueIndex;not null" json:"uuid"`
	Content string `gorm:"type:text;not null" json:"content"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
