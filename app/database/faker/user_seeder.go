package faker

import (
	"GoAppOne/app/models"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		Name:  faker.Name(),
		Email: faker.Email(),
	}
}
