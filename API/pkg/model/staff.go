package model

type Staff struct {
	Username     *string `json:"username"`
	Password     *string `json:"password"`
	HospitalCode *string `json:"hospitalCode"`
}
