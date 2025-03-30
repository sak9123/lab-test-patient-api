package entity

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID           uint    `gorm:"primarykey"`
	FirstNameTH  *string `gorm:"type:VARCHAR(200)"`
	MiddleNameTH *string `gorm:"type:VARCHAR(200)"`
	LastNameTH   *string `gorm:"type:VARCHAR(200)"`
	FirstNameEN  *string `gorm:"type:VARCHAR(200)"`
	MiddleNameEN *string `gorm:"type:VARCHAR(200)"`
	LastNameEN   *string `gorm:"type:VARCHAR(200)"`
	DateOfBirth  *time.Time
	PatientHN    *string `gorm:"type:VARCHAR(10)"`
	NationalId   *string `gorm:"type:VARCHAR(13)"`
	PassportId   *string `gorm:"type:VARCHAR(20)"`
	PhoneNumber  *string `gorm:"type:VARCHAR(25)"`
	Email        *string `gorm:"type:VARCHAR(100)"`
	Gender       *string `gorm:"type:CHAR(1)"`
	HospitalCode *string `gorm:"type:CHAR(2)"`
	CreatedBy    *string `gorm:"type:VARCHAR(100)"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	UpdatedBy    *string `gorm:"type:VARCHAR(100)"`
}

func (Patient) TableName() string {
	return "patient"
}

func PatientSeeds(db *gorm.DB) {
	createAt := time.Now()
	createdBy := "system"

	patients := []Patient{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 10; i++ {
		genders := []string{"M", "F"}
		gender := genders[r.Intn(len(genders))]
		firstNameTH := "firstNameTH" + strconv.Itoa(i)
		middleNameTH := "firstNameTH" + strconv.Itoa(i)
		lastNameTH := "lastNameTH" + strconv.Itoa(i)
		firstNameEN := "firstNameEN" + strconv.Itoa(i)
		middleNameEN := "middleNameEN" + strconv.Itoa(i)
		lastNameEN := "lastNameEN" + strconv.Itoa(i)
		patientHN := "patientHN" + strconv.Itoa(i)
		phoneNumber := "phoneNumber"
		email := "email" + strconv.Itoa(i)
		hospitalCode := "01"
		num := 1000000000000 + r.Int63n(9000000000000)
		nationalId := strconv.FormatInt(num, 10)
		num = 1000000000000 + r.Int63n(9000000000000)
		passportId := strconv.FormatInt(num, 10)
		dateOfBirth := time.Date(1995, time.March, 1, 0, 0, 0, 0, time.UTC)
		patients = append(patients, Patient{
			FirstNameTH:  &firstNameTH,
			MiddleNameTH: &middleNameTH,
			LastNameTH:   &lastNameTH,
			FirstNameEN:  &firstNameEN,
			MiddleNameEN: &middleNameEN,
			LastNameEN:   &lastNameEN,
			DateOfBirth:  &dateOfBirth,
			PatientHN:    &patientHN,
			NationalId:   &nationalId,
			PassportId:   &passportId,
			PhoneNumber:  &phoneNumber,
			Email:        &email,
			Gender:       &gender,
			HospitalCode: &hospitalCode,
			CreatedBy:    &createdBy,
			CreatedAt:    &createAt,
			UpdatedBy:    &createdBy,
			UpdatedAt:    &createAt,
		})
	}

	for _, entity := range patients {
		err := db.Save(&entity).Error
		if err != nil {
			fmt.Printf("Error when create patient: %s\n", *entity.FirstNameTH+" "+*entity.LastNameTH)
		} else {
			fmt.Printf("Success create patient: %s\n", *entity.FirstNameTH+" "+*entity.LastNameTH)
		}
	}
}
