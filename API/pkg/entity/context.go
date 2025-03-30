package entity

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	log.Println("waiting autoMigrate..")

	seedData(db)
	err := db.AutoMigrate(
		&Hospital{},
		&Patient{},
		&Staff{},
	)
	if err != nil {
		panic(err)
	}

	log.Println("autoMigrate successfully..")
}

func seedData(db *gorm.DB) {
	var errDb error
	if errDb = db.AutoMigrate(&Hospital{}); errDb == nil && db.Migrator().HasTable(&Hospital{}) {
		if err := db.First(&Hospital{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			HospitalSeeds(db)
		}
	}

	if errDb != nil {
		panic(errDb)
	}

	if errDb = db.AutoMigrate(&Patient{}); errDb == nil && db.Migrator().HasTable(&Patient{}) {
		if err := db.First(&Patient{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			PatientSeeds(db)
		}
	}

	if errDb != nil {
		panic(errDb)
	}

}
