package model

type StaffCriteria struct {
	Username     *string `json:"username"`
	Password     *string `json:"password"`
	HospitalCode *string `json:"hospitalCode"`
}
