package service

import (
	"hospitalApi/cmd/config"
	"hospitalApi/pkg/entity"
	"hospitalApi/pkg/errs"
	"hospitalApi/pkg/helper"
	"hospitalApi/pkg/model"
	"hospitalApi/pkg/repository"
)

type IStaffService interface {
	Create(input model.Staff) (isSuccess bool, err *errs.Error)
	IsExistsUsername(username string) (err *errs.Error)
	Login(input model.StaffCriteria) (result string, err *errs.Error)
	ValidateSave(input model.Staff) (err *errs.Error)
	ValidateLogin(input model.StaffCriteria) (err *errs.Error)
}

type StaffService struct {
	UseCase          IStaffService
	Common           helper.ICommon
	StaffsRepository repository.IStaffsRepository
}

func MakeIStaffService(
	iCommon helper.ICommon,
	StaffsRepository repository.IStaffsRepository,
) IStaffService {
	StaffService := &StaffService{
		Common:           iCommon,
		StaffsRepository: StaffsRepository,
	}

	StaffService.UseCase = StaffService
	return StaffService
}

func (s *StaffService) Create(input model.Staff) (isSuccess bool, err *errs.Error) {
	defer handlePanic("Create", &err)

	serviceErr := s.UseCase.ValidateSave(input)
	if serviceErr != nil {
		return false, serviceErr
	}

	serviceErr = s.UseCase.IsExistsUsername(*input.Username)
	if serviceErr != nil {
		return false, serviceErr
	}

	password, _ := s.Common.HashPassword(*input.Password)
	entity := entity.Staff{
		Username:     *input.Username,
		Password:     password,
		HospitalCode: *input.HospitalCode,
	}

	repoErr := s.StaffsRepository.Create(entity)
	if repoErr != nil {
		return false, errs.NewInternalServerError(repoErr.Error())
	}

	return true, nil
}

func (s *StaffService) ValidateSave(input model.Staff) (err *errs.Error) {
	if input.Username == nil {
		return errs.NewNotImplementedError("please username")
	}

	if input.Password == nil {
		return errs.NewNotImplementedError("please password")
	}

	if input.HospitalCode == nil {
		return errs.NewNotImplementedError("please hospital code")
	}

	return nil
}

func (s *StaffService) IsExistsUsername(username string) (err *errs.Error) {
	staffs, repoErr := s.StaffsRepository.Get(model.StaffCriteria{Username: &username})
	if repoErr != nil {
		return errs.NewInternalServerError(repoErr.Error())
	}

	if len(staffs) > 0 {
		return errs.NewUnprocessableEntityError("That username already exists on the system.")
	}

	return nil
}

func (s *StaffService) ValidateLogin(input model.StaffCriteria) (err *errs.Error) {
	if input.Username == nil {
		return errs.NewNotImplementedError("please username")
	}

	if input.Password == nil {
		return errs.NewNotImplementedError("please password")
	}

	if input.HospitalCode == nil {
		return errs.NewNotImplementedError("please hospital code")
	}

	return nil
}

func (s *StaffService) Login(input model.StaffCriteria) (result string, err *errs.Error) {
	defer handlePanic("Get", &err)

	serviceErr := s.UseCase.ValidateLogin(input)
	if serviceErr != nil {
		return "", serviceErr
	}

	staffs, repoErr := s.StaffsRepository.Get(input)
	if repoErr != nil {
		return "", errs.NewInternalServerError(repoErr.Error())
	}

	if len(staffs) == 0 {
		return "", errs.NewNotFoundError("Invalid username or password.")
	}

	if !s.Common.ValidPassword(*input.Password) {
		return "", errs.NewNotFoundError("Invalid username or password.")
	}

	configuration := config.New()
	token, commonErr := s.Common.GenerateToken(*input.Username, configuration.SecretKey)
	if commonErr != nil {
		return "", errs.NewInternalServerError(commonErr.Error())
	}

	return token, nil
}
