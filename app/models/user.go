package models

type User struct {
	ID    uint
	Name  string
	Email string `gorm:"uniqueIndex;not null"`
	Books []Book `gorm:"foreignKey:AuthorID"`
}
