package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    int    `gorm:"primarykey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"uniqueIndex; not null"`
	Age   int    `gorm:"not null"`
	Posts []Post `gorm:"foreignkey:AuthorID"`
}

type Post struct {
	gorm.Model
	ID       int    `gorm:"primarykey"`
	Title    string `gorm:"not null"`
	Content  string `gorm:"type:text;not null"`
	AuthorID int    `gorm:"not null"`
	Author   User   `gorm:"foreignkey:AuthorID"`
}
