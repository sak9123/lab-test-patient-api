package handler

import (
	"encoding/json"
	"hospitalApi/pkg/errs"
	"hospitalApi/pkg/helper"
	"hospitalApi/pkg/model"
	"hospitalApi/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	Common         helper.ICommon
	PatientService service.IPatientService
}

func MakePatientHandler(iCommon *helper.ICommon,
	IPatientService *service.IPatientService) *PatientHandler {
	return &PatientHandler{
		Common:         *iCommon,
		PatientService: *IPatientService,
	}
}

func (h PatientHandler) Search(c *gin.Context) {
	w := c.Writer
	r := c.Request

	var criteria model.PatientCriteria
	validateErr := json.NewDecoder(r.Body).Decode(&criteria)
	if validateErr != nil {
		err := errs.NewBadRequestError(ErrRequestBodyIsNotValid)
		h.Common.HandleErr(&w, err)
		return
	}

	criteria.Username = GetLoginUserFromRequest(r)
	result, err := h.PatientService.Get(criteria)
	if err != nil {
		h.Common.HandleErr(&w, err)
	}

	h.Common.APIResponse(&w, http.StatusOK, result)
}

func (h PatientHandler) SearchById(c *gin.Context) {
	w := c.Writer
	r := c.Request
	id := c.Param("id")
	criteria := model.PatientCriteria{
		NationalId: &id,
		PassportId: &id,
	}

	criteria.Username = GetLoginUserFromRequest(r)
	result, err := h.PatientService.Get(criteria)
	if err != nil {
		h.Common.HandleErr(&w, err)
	}

	h.Common.APIResponse(&w, http.StatusOK, result)
}
