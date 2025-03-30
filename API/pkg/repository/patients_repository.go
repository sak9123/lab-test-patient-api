package repository

import (
	"errors"
	"hospitalApi/pkg/entity"
	"hospitalApi/pkg/model"

	"gorm.io/gorm"
)

type IPatientsRepository interface {
	Get(input model.PatientCriteria) (result []entity.Patient, err error)
}

type PatientsRepository struct {
	DB *gorm.DB
}

func MakeIPatientsRepository(db *gorm.DB) IPatientsRepository {
	return &PatientsRepository{
		DB: db,
	}
}

func (r PatientsRepository) Where(tx *gorm.DB, input model.PatientCriteria) (result *gorm.DB, err error) {

	if input.Username != nil {
		tx = tx.Where(" EXISTS (SELECT 1 FROM staff WHERE username =?)", *input.Username)
	}

	if input.NationalId != nil {
		tx = tx.Where("national_id = ?", *input.NationalId)
	}

	if input.PassportId != nil {
		tx = tx.Or("passport_id = ?", *input.PassportId)
	}

	if input.FirstName != nil {
		firstName := "%" + *input.FirstName + "%"
		tx = tx.Or("first_name_th LIKE ?", firstName)
		tx = tx.Or("first_name_en LIKE ?", firstName)
	}

	if input.LastName != nil {
		lastName := "%" + *input.LastName + "%"
		tx = tx.Or("last_name_th LIKE ?", lastName)
		tx = tx.Or("last_name_en LIKE ?", lastName)
	}

	if input.DateOfBirth != nil {
		dateOfBirth := DateToSQLString(*input.DateOfBirth)
		tx = tx.Or("date_Of_birth  = ?", dateOfBirth)
	}

	if input.PhoneNumber != nil {
		tx = tx.Or("phone_number LIKE ?", *input.PhoneNumber)
	}

	return tx, nil
}

func (r PatientsRepository) Get(input model.PatientCriteria) (result []entity.Patient, err error) {
	tx := r.DB

	tx, err = r.Where(tx, input)
	if err != nil {
		return nil, errors.New("Patient " + err.Error())
	}

	tx = tx.Find(&result)
	if tx.Error != nil {
		return nil, errors.New("Patient" + tx.Error.Error())
	}

	return result, nil
}
