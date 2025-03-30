package entity

import (
	"time"
)

type Staff struct {
	ID           uint     `gorm:"primarykey"`
	Username     string   `gorm:"not null;type:VARCHAR(100)"`
	Password     string   `gorm:"not null;type:VARCHAR(2000)"`
	HospitalCode string   `gorm:"not null;type:VARCHAR(2)"`
	Hospital     Hospital `gorm:"foreignKey:HospitalCode;references:ID"`
	CreatedBy    *string  `gorm:"type:VARCHAR(100)"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	UpdatedBy    *string `gorm:"type:VARCHAR(100)"`
}

func (Staff) TableName() string {
	return "staff"
}

// func StaffSeeds(db *gorm.DB) {
// 	createAt := time.Now()
// 	createdBy := "system"
// 	password, _ := helper.Common.HashPassword("1234", "xxx")
// 	staffs := []Staff{}
// 	for i := 1; i <= 2; i++ {
// 		username := "username" + strconv.Itoa(i)
// 		staffs = append(staffs, Staff{
// 			Username:  username,
// 			Password:  password,
// 			CreatedBy: &createdBy,
// 			CreatedAt: &createAt,
// 			UpdatedBy: &createdBy,
// 			UpdatedAt: &createAt,
// 		})
// 	}

// 	for _, entity := range staffs {
// 		err := db.Save(&entity).Error
// 		if err != nil {
// 			fmt.Printf("Error when create staff: %s\n", entity.Username)
// 		} else {
// 			fmt.Printf("Success create staff: %s\n", entity.Username)
// 		}
// 	}
// }
