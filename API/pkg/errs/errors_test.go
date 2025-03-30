package errs_test

import (
	goErr "errors"
	"hospitalApi/pkg/errs"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	t.Run("should return error message", func(t *testing.T) {
		expected := "Error"
		err := errs.Error{
			Err: goErr.New(expected),
		}

		result := err.Error()

		assert.Equal(t, expected, result)
	})
}

func TestNewInternalServerError(t *testing.T) {
	t.Run("should return Error struct", func(t *testing.T) {
		expected := errs.Error{
			StatusCode: http.StatusInternalServerError,
			Err:        goErr.New("error"),
		}

		result := errs.NewInternalServerError(expected.Err.Error())

		assert.Equal(t, expected, *result)
	})
}

func TestNewNotImplementedError(t *testing.T) {
	t.Run("should return Error struct", func(t *testing.T) {
		expected := errs.Error{
			StatusCode: http.StatusNotImplemented,
			Err:        goErr.New("error"),
		}

		result := errs.NewNotImplementedError(expected.Err.Error())

		assert.Equal(t, expected, *result)
	})
}

func TestNewNotFoundError(t *testing.T) {
	t.Run("should return Error struct", func(t *testing.T) {
		expected := errs.Error{
			StatusCode: http.StatusNotFound,
			Err:        goErr.New("error"),
		}

		result := errs.NewNotFoundError(expected.Err.Error())

		assert.Equal(t, expected, *result)
	})
}

func TestNewUnauthorizedError(t *testing.T) {
	t.Run("should return Error struct", func(t *testing.T) {
		expected := errs.Error{
			StatusCode: http.StatusUnauthorized,
			Err:        goErr.New("error"),
		}

		result := errs.NewUnauthorizedError(expected.Err.Error())

		assert.Equal(t, expected, *result)
	})
}

func TestNewBadRequestError(t *testing.T) {
	t.Run("should return Error struct", func(t *testing.T) {
		expected := errs.Error{
			StatusCode: http.StatusBadRequest,
			Err:        goErr.New("error"),
		}

		result := errs.NewBadRequestError(expected.Err.Error())

		assert.Equal(t, expected, *result)
	})
}

func TestNewUnprocessableEntityError(t *testing.T) {
	t.Run("should return Error struct", func(t *testing.T) {
		expected := errs.Error{
			StatusCode: http.StatusUnprocessableEntity,
			Err:        goErr.New("error"),
		}

		result := errs.NewUnprocessableEntityError(expected.Err.Error())

		assert.Equal(t, expected, *result)
	})
}

func TestNNewConflictError(t *testing.T) {
	t.Run("should return Error struct", func(t *testing.T) {
		expected := errs.Error{
			StatusCode: http.StatusConflict,
			Err:        goErr.New("error"),
		}

		result := errs.NewConflictError(expected.Err.Error())

		assert.Equal(t, expected, *result)
	})
}
