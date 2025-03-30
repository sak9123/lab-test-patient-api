package repository

import (
	"errors"
	"hospitalApi/pkg/entity"
	"hospitalApi/pkg/helper"
	"hospitalApi/pkg/model"

	"gorm.io/gorm"
)

type IStaffsRepository interface {
	Create(input entity.Staff) (err error)
	Get(input model.StaffCriteria) (result []entity.Staff, err error)
}

type StaffsRepository struct {
	Common helper.ICommon
	DB     *gorm.DB
}

func MakeIStaffsRepository(db *gorm.DB, iCommon helper.ICommon) IStaffsRepository {
	return &StaffsRepository{
		DB:     db,
		Common: iCommon,
	}
}

func (r StaffsRepository) Where(tx *gorm.DB, input model.StaffCriteria) (result *gorm.DB, err error) {
	if input.Username != nil {
		tx = tx.Where("username = ?", *input.Username)
	}

	if input.HospitalCode != nil {
		tx = tx.Where("hospital_code = ?", *input.HospitalCode)
	}

	return tx, nil
}

func (r StaffsRepository) Get(input model.StaffCriteria) (result []entity.Staff, err error) {
	tx := r.DB

	tx, err = r.Where(tx, input)
	if err != nil {
		return nil, errors.New("Staff " + err.Error())
	}

	tx = tx.Find(&result)
	if tx.Error != nil {
		return nil, errors.New("Staff" + tx.Error.Error())
	}

	return result, nil
}

func (r StaffsRepository) Create(input entity.Staff) (err error) {
	tx := r.DB.Create(&input)
	if tx.Error != nil {
		return errors.New("Staff " + tx.Error.Error())
	}

	return nil
}
