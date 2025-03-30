package model

import "time"

type Patient struct {
	ID           *uint      `json:"id"`
	FirstNameTH  *string    `json:"firstNameTH"`
	MiddleNameTH *string    `json:"middleNameTH"`
	LastNameTH   *string    `json:"lastNameTH"`
	FirstNameEN  *string    `json:"firstNameEN"`
	MiddleNameEN *string    `json:"middleNameEN"`
	LastNameEN   *string    `json:"lastNameEN"`
	DateOfBirth  *time.Time `json:"dateOfBirth"`
	PatientHN    *string    `json:"patientHN"`
	NationalId   *string    `json:"nationalId"`
	PassportId   *string    `json:"passportId"`
	PhoneNumber  *string    `json:"phoneNumber"`
	Email        *string    `json:"email"`
	Gender       *string    `json:"gender"`
	HospitalCode *string    `json:"hospitalCode"`
}
