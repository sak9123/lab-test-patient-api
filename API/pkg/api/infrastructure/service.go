package infrastructure

import (
	"hospitalApi/pkg/service"
)

type Service struct {
	IPatientService service.IPatientService
	IStaffService   service.IStaffService
}

func CreateService(repo Repository, helper Helper) (result Service) {
	result = Service{}
	result.IPatientService = service.MakeIPatientService(
		helper.ICommon,
		repo.IPatientsRepository,
	)

	result.IStaffService = service.MakeIStaffService(
		helper.ICommon,
		repo.IStaffsRepository,
	)

	return result
}
