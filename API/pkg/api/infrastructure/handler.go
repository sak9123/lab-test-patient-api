package infrastructure

import (
	"hospitalApi/pkg/handler"
)

type Handler struct {
	PatientHandler *handler.PatientHandler
	StaffHandler   *handler.StaffHandler
}

func CreateHandler(service Service, helper Helper) (result Handler) {
	return Handler{
		PatientHandler: handler.MakePatientHandler(&helper.ICommon, &service.IPatientService),
		StaffHandler:   handler.MakeStaffHandler(&helper.ICommon, &service.IStaffService),
	}
}
