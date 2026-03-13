package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	UUID    string `gorm:"uniqueIndex;not null" json:"uuid"`
	Content string `gorm:"type:text;not null" json:"content"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
