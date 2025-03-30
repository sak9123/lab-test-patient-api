package service

import (
	"hospitalApi/pkg/entity"
	"hospitalApi/pkg/errs"
	"hospitalApi/pkg/helper"
	"hospitalApi/pkg/model"
	"hospitalApi/pkg/repository"
)

type IPatientService interface {
	Get(input model.PatientCriteria) (result []model.Patient, err *errs.Error)
	MappingEntityToModel(input entity.Patient) (result *model.Patient)
}

type PatientService struct {
	UseCase            IPatientService
	Common             helper.ICommon
	PatientsRepository repository.IPatientsRepository
}

func MakeIPatientService(
	iCommon helper.ICommon,
	patientsRepository repository.IPatientsRepository,
) IPatientService {
	patientService := &PatientService{
		Common:             iCommon,
		PatientsRepository: patientsRepository,
	}

	patientService.UseCase = patientService
	return patientService
}

func (s *PatientService) Get(input model.PatientCriteria) (result []model.Patient, err *errs.Error) {
	defer handlePanic("Get", &err)

	patients, repoErr := s.PatientsRepository.Get(input)
	if repoErr != nil {
		return nil, errs.NewInternalServerError(repoErr.Error())
	}

	for _, entity := range patients {
		resultMapping := s.UseCase.MappingEntityToModel(entity)
		result = append(result, *resultMapping)
	}

	return result, nil
}

func (s *PatientService) MappingEntityToModel(input entity.Patient) (result *model.Patient) {
	result = &model.Patient{
		ID:           &input.ID,
		FirstNameTH:  input.FirstNameTH,
		MiddleNameTH: input.MiddleNameTH,
		LastNameTH:   input.LastNameTH,
		FirstNameEN:  input.FirstNameEN,
		MiddleNameEN: input.MiddleNameEN,
		LastNameEN:   input.LastNameEN,
		DateOfBirth:  input.DateOfBirth,
		PatientHN:    input.PatientHN,
		NationalId:   input.NationalId,
		PassportId:   input.PassportId,
		PhoneNumber:  input.PhoneNumber,
		Email:        input.Email,
		Gender:       input.Gender,
		HospitalCode: input.HospitalCode,
	}

	return result
}
