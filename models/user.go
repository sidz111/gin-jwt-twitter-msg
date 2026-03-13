package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UUID     string `gorm:"type:char(36);uniqueIndex;not null" json:"uuid"`
	Username string `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Email    string `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Bio      string `json:"bio"`
	// Token    string `json:"token"`

	Posts []Post `gorm:"foreignKey:UserID" json:"posts"`
}
