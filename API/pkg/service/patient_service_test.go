package service_test

import (
	"errors"
	"hospitalApi/pkg/entity"
	"hospitalApi/pkg/errs"
	"hospitalApi/pkg/helper"
	"hospitalApi/pkg/mocks"
	"hospitalApi/pkg/model"
	"hospitalApi/pkg/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMakePatientService(t *testing.T) {
	result := service.MakeIPatientService(
		&helper.Common{},
		&mocks.IPatientsRepository{},
	)
	assert.NotNil(t, result)
}

func TestGet(t *testing.T) {
	var (
		common                     *mocks.ICommon
		useCase                    *mocks.IPatientService
		patientService             *service.PatientService
		patientsRepository         *mocks.IPatientsRepository
		resultGet                  []entity.Patient
		resultGetErr               error
		resultMappingEntityToModel *model.Patient
		input                      model.PatientCriteria
	)

	beforeEach := func() {
		common = &mocks.ICommon{}
		useCase = &mocks.IPatientService{}
		patientsRepository = &mocks.IPatientsRepository{}
		patientService = &service.PatientService{
			Common:             common,
			UseCase:            useCase,
			PatientsRepository: patientsRepository,
		}

		resultGetErr = nil
		nationalId := "nationalId"
		input = model.PatientCriteria{
			NationalId: &nationalId,
		}

		resultGet = []entity.Patient{{}, {}}
		resultMappingEntityToModel = &model.Patient{}
		patientsRepository.On("Get", mock.Anything).Return(
			func(input model.PatientCriteria) (result []entity.Patient, err error) {
				return resultGet, resultGetErr
			},
		)

		useCase.On("MappingEntityToModel", mock.Anything).Return(
			func(input entity.Patient) (result *model.Patient) {
				return resultMappingEntityToModel
			},
		)

	}

	t.Run("should return server error when get  was fail", func(t *testing.T) {
		beforeEach()
		resultGetErr = errors.New("error create")
		expected := errs.NewInternalServerError(resultGetErr.Error())

		result, err := patientService.Get(input)

		assert.Nil(t, result)
		assert.Equal(t, expected, err)
	})

	t.Run("should get patient by criteria", func(t *testing.T) {
		beforeEach()
		expected := input

		patientService.Get(input)

		patientsRepository.AssertCalled(t, "Get", expected)
	})

	t.Run("should mapping entity to model", func(t *testing.T) {
		beforeEach()
		expected1 := resultGet[0]
		expected2 := resultGet[1]

		patientService.Get(input)

		useCase.AssertNumberOfCalls(t, "MappingEntityToModel", 2)
		useCase.AssertCalled(t, "MappingEntityToModel", expected1)
		useCase.AssertCalled(t, "MappingEntityToModel", expected2)
	})

	t.Run("should return result", func(t *testing.T) {
		beforeEach()
		expected := []model.Patient{{}, {}}

		result, err := patientService.Get(input)

		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})

}

func TestMappingEntityToModel(t *testing.T) {
	var (
		patientService *service.PatientService
		input          entity.Patient
	)

	beforeEach := func() {
		patientService = &service.PatientService{}
		input = entity.Patient{
			ID:           1,
			FirstNameTH:  helper.NewString("FirstNameTH"),
			MiddleNameTH: helper.NewString("MiddleNameTH"),
			LastNameTH:   helper.NewString("LastNameTH"),
			FirstNameEN:  helper.NewString("FirstNameEN"),
			MiddleNameEN: helper.NewString("MiddleNameEN"),
			LastNameEN:   helper.NewString("LastNameEN"),
			DateOfBirth:  helper.NewTime(time.Now()),
			PatientHN:    helper.NewString("PatientHN"),
			NationalId:   helper.NewString("NationalId"),
			PassportId:   helper.NewString("PassportId"),
			PhoneNumber:  helper.NewString("PhoneNumber"),
			Email:        helper.NewString("Email"),
			Gender:       helper.NewString("Gender"),
			HospitalCode: helper.NewString("HospitalCode"),
		}
	}

	t.Run("should return model patient", func(t *testing.T) {
		beforeEach()
		expected := &model.Patient{
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

		result := patientService.MappingEntityToModel(input)

		assert.Equal(t, expected, result)
	})

}
