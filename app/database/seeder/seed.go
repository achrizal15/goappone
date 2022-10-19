package seeder

import (
	"GoAppOne/app/database/faker"

	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: faker.UserFaker(db)},
		{Seeder: faker.BookFaker(db)},
	}
}
func RunSeeder(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}
	return nil
}
