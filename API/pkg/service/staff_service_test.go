package service_test

import (
	"errors"
	"hospitalApi/cmd/config"
	"hospitalApi/pkg/entity"
	"hospitalApi/pkg/errs"
	"hospitalApi/pkg/helper"
	"hospitalApi/pkg/mocks"
	"hospitalApi/pkg/model"
	"hospitalApi/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMakeStaffService(t *testing.T) {
	result := service.MakeIStaffService(
		&helper.Common{},
		&mocks.IStaffsRepository{},
	)
	assert.NotNil(t, result)
}

func TestCreate(t *testing.T) {
	var (
		common                    *mocks.ICommon
		useCase                   *mocks.IStaffService
		staffService              *service.StaffService
		staffsRepository          *mocks.IStaffsRepository
		resultValidateSaveErr     *errs.Error
		resultHashPassword        string
		resultHashPasswordErr     error
		resultIsExistsUsernameErr *errs.Error
		resultCreateErr           error
		input                     model.Staff
	)

	beforeEach := func() {
		common = &mocks.ICommon{}
		useCase = &mocks.IStaffService{}
		staffsRepository = &mocks.IStaffsRepository{}
		staffService = &service.StaffService{
			Common:           common,
			UseCase:          useCase,
			StaffsRepository: staffsRepository,
		}

		resultValidateSaveErr = nil
		resultHashPasswordErr = nil
		resultIsExistsUsernameErr = nil
		resultCreateErr = nil

		resultHashPassword = "password"

		username := "username"
		password := "password"
		hospitalCode := "hospitalCode"
		input = model.Staff{
			Username:     &username,
			Password:     &password,
			HospitalCode: &hospitalCode,
		}

		useCase.On("ValidateSave", mock.Anything).Return(
			func(input model.Staff) (err *errs.Error) {
				return resultValidateSaveErr
			},
		)

		useCase.On("IsExistsUsername", mock.Anything, mock.Anything).Return(
			func(username string, hospitalCode string) (err *errs.Error) {
				return resultIsExistsUsernameErr
			},
		)

		common.On("HashPassword", mock.Anything).Return(
			func(password string) (hasPassword string, err error) {
				return resultHashPassword, resultHashPasswordErr
			},
		)

		staffsRepository.On("Create", mock.Anything).Return(
			func(entity entity.Staff) (err error) {
				return resultCreateErr
			},
		)

	}

	t.Run("should return new not implemented error when validate save was invalid input", func(t *testing.T) {
		beforeEach()
		resultValidateSaveErr = errs.NewNotImplementedError("error validate save")
		expected := false

		result, err := staffService.Create(input)

		assert.Equal(t, resultValidateSaveErr, err)
		assert.Equal(t, expected, result)
	})

	t.Run("should call validate save", func(t *testing.T) {
		beforeEach()
		expected := input

		staffService.Create(input)

		useCase.AssertCalled(t, "ValidateSave", expected)
	})

	t.Run("should return new not implemented error when username is exists.", func(t *testing.T) {
		beforeEach()
		resultIsExistsUsernameErr = errs.NewNotImplementedError("error username is exists")
		expected := false

		result, err := staffService.Create(input)

		assert.Equal(t, resultIsExistsUsernameErr, err)
		assert.Equal(t, expected, result)
	})

	t.Run("should check exists user name", func(t *testing.T) {
		beforeEach()
		expectedUsername := *input.Username
		expectedHospitalCode := *input.HospitalCode

		staffService.Create(input)

		useCase.AssertCalled(t, "IsExistsUsername", expectedUsername, expectedHospitalCode)
	})

	t.Run("should hash password", func(t *testing.T) {
		beforeEach()
		expected := *input.Password

		staffService.Create(input)

		common.AssertCalled(t, "HashPassword", expected)
	})

	t.Run("should return server error when Create was fail", func(t *testing.T) {
		beforeEach()
		resultCreateErr = errors.New("error create")
		expectedErr := errs.NewInternalServerError(resultCreateErr.Error())
		expected := false

		result, err := staffService.Create(input)

		assert.Equal(t, expectedErr, err)
		assert.Equal(t, expected, result)
	})

	t.Run("should create staff", func(t *testing.T) {
		beforeEach()
		expected := true

		result, err := staffService.Create(input)

		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})

}

func TestValidateSave(t *testing.T) {
	var (
		staffService *service.StaffService
		input        model.Staff
	)

	beforeEach := func() {
		staffService = &service.StaffService{}
		username := "username"
		password := "password"
		hospitalCode := "hospitalCode"
		input = model.Staff{
			Username:     &username,
			Password:     &password,
			HospitalCode: &hospitalCode,
		}
	}

	t.Run("should return new not implemented error when input username is nil", func(t *testing.T) {
		beforeEach()
		input.Username = nil
		expected := errs.NewNotImplementedError("please username")

		err := staffService.ValidateSave(input)

		assert.Equal(t, expected, err)
	})

	t.Run("should return new not implemented error when input password is nil", func(t *testing.T) {
		beforeEach()
		input.Password = nil
		expected := errs.NewNotImplementedError("please password")

		err := staffService.ValidateSave(input)

		assert.Equal(t, expected, err)
	})

	t.Run("should return new not implemented error when input hospital code is nil", func(t *testing.T) {
		beforeEach()
		input.HospitalCode = nil
		expected := errs.NewNotImplementedError("please hospital code")

		err := staffService.ValidateSave(input)

		assert.Equal(t, expected, err)
	})

}

func TestValidateLogin(t *testing.T) {
	var (
		staffService *service.StaffService
		input        model.StaffCriteria
	)

	beforeEach := func() {
		staffService = &service.StaffService{}
		username := "username"
		password := "password"
		hospitalCode := "hospitalCode"
		input = model.StaffCriteria{
			Username:     &username,
			Password:     &password,
			HospitalCode: &hospitalCode,
		}
	}

	t.Run("should return new not implemented error when input username is nil", func(t *testing.T) {
		beforeEach()
		input.Username = nil
		expected := errs.NewNotImplementedError("please username")

		err := staffService.ValidateLogin(input)

		assert.Equal(t, expected, err)
	})

	t.Run("should return new not implemented error when input password is nil", func(t *testing.T) {
		beforeEach()
		input.Password = nil
		expected := errs.NewNotImplementedError("please password")

		err := staffService.ValidateLogin(input)

		assert.Equal(t, expected, err)
	})

	t.Run("should return new not implemented error when input hospital code is nil", func(t *testing.T) {
		beforeEach()
		input.HospitalCode = nil
		expected := errs.NewNotImplementedError("please hospital code")

		err := staffService.ValidateLogin(input)

		assert.Equal(t, expected, err)
	})

}

func TestIsExistsUsername(t *testing.T) {
	var (
		staffsRepository  *mocks.IStaffsRepository
		staffService      *service.StaffService
		resultGet         []entity.Staff
		resultGetErr      error
		inputUsername     string
		inputHospitalCode string
	)

	beforeEach := func() {
		staffsRepository = &mocks.IStaffsRepository{}
		staffService = &service.StaffService{
			StaffsRepository: staffsRepository,
		}
		inputUsername = "username"
		inputHospitalCode = "hospitalCode"

		resultGetErr = nil
		resultGet = []entity.Staff{}

		staffsRepository.On("Get", mock.Anything).Return(
			func(input model.StaffCriteria) (result []entity.Staff, err error) {
				return resultGet, resultGetErr
			},
		)
	}

	t.Run("should return server error when get staff was fail", func(t *testing.T) {
		beforeEach()
		resultGetErr = errors.New("error get staff")
		expected := errs.NewInternalServerError(resultGetErr.Error())

		err := staffService.IsExistsUsername(inputUsername, inputHospitalCode)

		assert.Equal(t, expected, err)
	})

	t.Run("should get staff by user name", func(t *testing.T) {
		beforeEach()
		expected := model.StaffCriteria{Username: &inputUsername, HospitalCode: &inputHospitalCode}

		staffService.IsExistsUsername(inputUsername, inputHospitalCode)

		staffsRepository.AssertCalled(t, "Get", expected)
	})

	t.Run("should return new un processable entity error when username is exists.", func(t *testing.T) {
		beforeEach()
		resultGet = []entity.Staff{{}}
		expected := errs.NewUnprocessableEntityError("That username already exists on the system.")

		err := staffService.IsExistsUsername(inputUsername, inputHospitalCode)

		assert.Equal(t, expected, err)
	})

	t.Run("should return nil when not have user in system", func(t *testing.T) {
		beforeEach()
		resultGet = []entity.Staff{}

		err := staffService.IsExistsUsername(inputUsername, inputHospitalCode)

		assert.Nil(t, err)
	})

}

func TestLogin(t *testing.T) {
	var (
		common                 *mocks.ICommon
		useCase                *mocks.IStaffService
		staffService           *service.StaffService
		staffsRepository       *mocks.IStaffsRepository
		resultValidateLoginErr *errs.Error
		resultGenerateToken    string
		resultGenerateTokenErr error
		resultValidPassword    bool
		resultGet              []entity.Staff
		resultGetErr           error
		input                  model.StaffCriteria
	)

	beforeEach := func() {
		common = &mocks.ICommon{}
		useCase = &mocks.IStaffService{}
		staffsRepository = &mocks.IStaffsRepository{}
		staffService = &service.StaffService{
			Common:           common,
			UseCase:          useCase,
			StaffsRepository: staffsRepository,
		}

		resultValidateLoginErr = nil
		resultGenerateTokenErr = nil
		resultValidPassword = true
		resultGetErr = nil

		resultGenerateToken = "token"
		resultGet = []entity.Staff{}

		username := "username"
		password := "password"
		hospitalCode := "hospitalCode"
		input = model.StaffCriteria{
			Username:     &username,
			Password:     &password,
			HospitalCode: &hospitalCode,
		}

		useCase.On("ValidateLogin", mock.Anything).Return(
			func(input model.StaffCriteria) (err *errs.Error) {
				return resultValidateLoginErr
			},
		)

		common.On("ValidPassword", mock.Anything).Return(
			func(password string) bool {
				return resultValidPassword
			},
		)

		common.On("GenerateToken", mock.Anything, mock.Anything).Return(
			func(username string, jwtSecret string) (token string, err error) {
				return resultGenerateToken, resultGenerateTokenErr
			},
		)

		staffsRepository.On("Get", mock.Anything).Return(
			func(input model.StaffCriteria) (result []entity.Staff, err error) {
				return resultGet, resultGetErr
			},
		)

	}

	t.Run("should return new not implemented error when validate login was invalid input", func(t *testing.T) {
		beforeEach()
		resultValidateLoginErr = errs.NewNotImplementedError("error invalid login")
		expectedErr := resultValidateLoginErr
		expectedResult := ""

		result, err := staffService.Login(input)

		assert.Equal(t, expectedResult, result)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should call validate login", func(t *testing.T) {
		beforeEach()
		expected := input

		staffService.Login(input)

		useCase.AssertCalled(t, "ValidateLogin", expected)
	})

	t.Run("should return server error when get staff was fail", func(t *testing.T) {
		beforeEach()
		resultGetErr = errors.New("error get staff")
		expectedErr := errs.NewInternalServerError(resultGetErr.Error())
		expectedResult := ""

		result, err := staffService.Login(input)

		assert.Equal(t, expectedResult, result)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should get staff by user name", func(t *testing.T) {
		beforeEach()
		expected := input

		staffService.Login(input)

		staffsRepository.AssertCalled(t, "Get", expected)
	})

	t.Run("should return new not found error when username not found.", func(t *testing.T) {
		beforeEach()
		resultGet = []entity.Staff{}
		expectedErr := errs.NewNotFoundError("Invalid username or password.")
		expectedResult := ""
		result, err := staffService.Login(input)

		assert.Equal(t, expectedResult, result)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return new not found error when password invalid.", func(t *testing.T) {
		beforeEach()
		resultValidPassword = false
		expectedErr := errs.NewNotFoundError("Invalid username or password.")
		expectedResult := ""
		result, err := staffService.Login(input)

		assert.Equal(t, expectedResult, result)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should generate token", func(t *testing.T) {
		beforeEach()
		resultGet = []entity.Staff{{}}
		expected := *input.Password

		staffService.Login(input)

		common.AssertCalled(t, "ValidPassword", expected)
	})

	t.Run("should return new not found error when generate token was  fail", func(t *testing.T) {
		beforeEach()
		resultGenerateTokenErr = errs.NewNotFoundError("Invalid username or password.")
		expectedErr := resultGenerateTokenErr
		expectedResult := ""
		result, err := staffService.Login(input)

		assert.Equal(t, expectedResult, result)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should generate token", func(t *testing.T) {
		beforeEach()
		resultGet = []entity.Staff{{}}
		configuration := config.New()
		expectedSecretKey := configuration.SecretKey
		expectedUsername := *input.Username

		staffService.Login(input)

		common.AssertCalled(t, "GenerateToken", expectedUsername, expectedSecretKey)
	})

	t.Run("should return new token when login is success.", func(t *testing.T) {
		beforeEach()
		resultGet = []entity.Staff{{}}
		expected := resultGenerateToken

		result, err := staffService.Login(input)

		assert.Nil(t, err)
		assert.Equal(t, expected, result)

	})
}
