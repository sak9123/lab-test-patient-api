package handler

import (
	"net/http"
)

func GetLoginUserFromRequest(r *http.Request) *string {
	username := r.Header.Get("username")
	return &username
}
