package infrastructure

import (
	"hospitalApi/pkg/helper"
)

type Helper struct {
	ICommon helper.ICommon
}

func CreateHelper() (result Helper) {
	helperCommon := Helper{
		ICommon: helper.MakeICommon(),
	}
	return helperCommon
}
