package faker

import (
	"GoAppOne/app/models"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func BookFaker(db *gorm.DB) *models.Book {
	return &models.Book{
		Name:     faker.ChineseName(),
		AuthorID: 1,
	}
}
