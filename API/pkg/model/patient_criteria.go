package model

import "time"

type PatientCriteria struct {
	NationalId  *string    `json:"nationalId"`
	PassportId  *string    `json:"passportId"`
	FirstName   *string    `json:"firstName"`
	MiddleName  *string    `json:"middleName"`
	LastName    *string    `json:"lastName"`
	DateOfBirth *time.Time `json:"dateOfBirth"`
	PhoneNumber *string    `json:"phoneNumber"`
	Email       *string    `json:"email"`
	Username    *string    `json:"username"`
}
