package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Hospital struct {
	ID        uint    `gorm:"primarykey"`
	Code      *string `gorm:"type:CHAR(2)"`
	Name      *string `gorm:"type:VARCHAR(200)"`
	Active    bool
	CreatedBy *string `gorm:"type:VARCHAR(100)"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	UpdatedBy *string `gorm:"type:VARCHAR(100)"`
}

func (Hospital) TableName() string {
	return "hospital"
}

func HospitalSeeds(db *gorm.DB) {
	createAt := time.Now()
	Code1 := "01"
	Code2 := "02"
	hospital1 := "Nakornthon"
	hospital2 := "Thainakarin"
	createdBy := "system"
	hospitals := []Hospital{{
		Code:      &Code1,
		Name:      &hospital1,
		Active:    true,
		CreatedAt: &createAt,
		UpdatedAt: &createAt,
		CreatedBy: &createdBy,
		UpdatedBy: &createdBy,
	}, {
		Code:      &Code2,
		Name:      &hospital2,
		Active:    true,
		CreatedAt: &createAt,
		UpdatedAt: &createAt,
		CreatedBy: &createdBy,
		UpdatedBy: &createdBy,
	}}

	for _, entity := range hospitals {
		err := db.Save(&entity).Error
		if err != nil {
			fmt.Printf("Error when create hospital: %s\n", *entity.Name)
		} else {
			fmt.Printf("Success create hospital: %s\n", *entity.Name)
		}
	}
}
