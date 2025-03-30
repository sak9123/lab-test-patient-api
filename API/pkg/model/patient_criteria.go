package model

import "time"

type PatientCriteria struct {
	NationalId  *string    `schema:"nationalId"`
	PassportId  *string    `schema:"passportId"`
	FirstName   *string    `schema:"firstName"`
	MiddleName  *string    `schema:"middleName"`
	LastName    *string    `schema:"lastName"`
	DateOfBirth *time.Time `schema:"dateOfBirth"`
	PhoneNumber *string    `schema:"phoneNumber"`
	Email       *string    `schema:"email"`
	Username    *string    `schema:"username"`
}
