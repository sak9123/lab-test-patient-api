package helper

import (
	"encoding/json"
	"hospitalApi/cmd/config"
	"hospitalApi/pkg/errs"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var CountRetried int = 10

type ICommon interface {
	APIResponse(w *gin.ResponseWriter, statusCode int, data interface{})
	GetConfiguration() config.Configuration
	GenerateToken(username string, jwtSecret string) (string, error)
	GetResponse(url string) (result []byte, err error)
	HashPassword(plainTextPassword string) (string, error)
	HandleErr(w *gin.ResponseWriter, err *errs.Error) errs.Error
	HandlePanic(r interface{}, w *gin.ResponseWriter) error
	NewBoolean(input bool) *bool
	NewFloat64(input float64) *float64
	NewInt(input int) *int
	NewString(input string) *string
	NewUInt(input uint) *uint
	ValidPassword(plainTextPassword string) bool
}

func MakeICommon() ICommon {
	common := &Common{}
	common.UseCase = common
	return common
}

type Common struct {
	UseCase ICommon
}

func (common Common) APIResponse(w *gin.ResponseWriter, statusCode int, data interface{}) {
	ApiResponse(w, statusCode, data)
}

func (common *Common) GetResponse(url string) (result []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (common Common) GetConfiguration() config.Configuration {
	return config.New()
}

func ApiResponse(w *gin.ResponseWriter, statusCode int, data interface{}) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(statusCode)
	_ = json.NewEncoder(*w).Encode(data)
}

func (common Common) HandleErr(w *gin.ResponseWriter, err *errs.Error) errs.Error {
	return HandleErr(w, err)
}

func HandleErr(w *gin.ResponseWriter, err *errs.Error) errs.Error {
	ApiResponse(w, err.StatusCode, err.Err.Error())
	return *err
}

func (common Common) HandlePanic(r interface{}, w *gin.ResponseWriter) error {
	if r != nil {
		errMsg := "Internal Server Exception: "
		switch x := r.(type) {
		case string:
			errMsg = errMsg + x
		case error:
			errMsg = errMsg + x.Error()
		default:
			errMsg = errMsg + "unexpected panic"
		}
		err := errs.NewInternalServerError(errMsg)
		return HandleErr(w, err)
	}

	return nil
}

func (common Common) TimeNow() *time.Time {
	timeNow := time.Now()
	return &timeNow
}

func (common Common) TimeUTCNow() *time.Time {
	utcTime := time.Now().UTC()
	return &utcTime
}

func (common Common) NewBoolean(input bool) *bool {
	return &input
}

func (common Common) NewUInt(input uint) *uint {
	return &input
}

func (common Common) NewInt(input int) *int {
	return &input
}

func (common Common) NewString(input string) *string {
	return &input
}

func (common Common) NewFloat64(input float64) *float64 {
	return &input
}

func (common Common) GenerateToken(username string, jwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	})

	jwtSecretBytes := []byte(jwtSecret)
	tokenString, err := token.SignedString(jwtSecretBytes)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (common Common) HashPassword(plainTextPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (common Common) ValidPassword(plainTextPassword string) bool {
	hash, _ := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcrypt.DefaultCost)
	err := bcrypt.CompareHashAndPassword(hash, []byte(plainTextPassword))
	return err == nil

}

func GetErrorMsgFromRecover(r interface{}) string {
	errMsg := "Internal Server Exception: "

	switch x := r.(type) {
	case string:
		errMsg = errMsg + x
	case error:
		errMsg = errMsg + x.Error()
	default:
		errMsg = errMsg + "unexpected panic"
	}

	return errMsg
}

func NewString(input string) *string {
	return &input
}

func NewTime(input time.Time) *time.Time {
	return &input
}

func NewBoolean(input bool) *bool {
	return &input
}

func NewInt(input int) *int {
	return &input
}

func NewInt64(input int64) *int64 {
	return &input
}

func NewUInt(input uint) *uint {
	return &input
}

func NewFloat64(input float64) *float64 {
	return &input
}
