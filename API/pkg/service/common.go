package service

import (
	"fmt"
	"hospitalApi/pkg/errs"
	"hospitalApi/pkg/helper"
	"runtime/debug"
)

func handlePanic(funcName string, err **errs.Error) {
	if r := recover(); r != nil {
		result := recoverPanic(r)
		*err = errs.NewInternalServerError("function name: '" + funcName + "' " + result.Error())
		return
	}
}

func recoverPanic(r interface{}) *errs.Error {
	fmt.Printf("Stack trace of panic: \n%s\n", string(debug.Stack()))
	msgErr := helper.GetErrorMsgFromRecover(r)
	return errs.NewInternalServerError(msgErr)
}
