package models

type Book struct {
	ID       uint
	Name     string `gorm:"not null;"`
	Author   User
	AuthorID uint `gorm:"not null;index"`
	// gorm.Model
}
