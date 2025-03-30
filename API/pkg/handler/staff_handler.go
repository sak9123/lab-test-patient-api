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

type StaffHandler struct {
	Common       helper.ICommon
	StaffService service.IStaffService
}

func MakeStaffHandler(iCommon *helper.ICommon,
	IStaffService *service.IStaffService) *StaffHandler {
	return &StaffHandler{
		Common:       *iCommon,
		StaffService: *IStaffService,
	}
}

func (h StaffHandler) Login(c *gin.Context) {
	w := c.Writer
	username := c.Query("username")
	password := c.Query("password")
	hospitalCode := c.Query("hospitalCode")

	criteria := model.StaffCriteria{
		Username:     &username,
		Password:     &password,
		HospitalCode: &hospitalCode,
	}
	token, err := h.StaffService.Login(criteria)
	if err != nil {
		h.Common.HandleErr(&w, err)
		return
	}

	h.Common.APIResponse(&w, http.StatusOK, token)
}

func (h StaffHandler) Create(c *gin.Context) {
	w := c.Writer
	r := c.Request

	var staff model.Staff
	validateErr := json.NewDecoder(r.Body).Decode(&staff)
	if validateErr != nil {
		err := errs.NewBadRequestError(ErrRequestBodyIsNotValid)
		h.Common.HandleErr(&w, err)
		return
	}

	token, err := h.StaffService.Create(staff)
	if err != nil {
		h.Common.HandleErr(&w, err)
		return
	}

	h.Common.APIResponse(&w, http.StatusOK, token)
}
