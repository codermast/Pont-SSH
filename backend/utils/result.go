package utils

import (
	"PontSsh/backend/constant"
	"PontSsh/backend/entity"
)

func Success(data interface{}, msg string) entity.Result {
	result := entity.Result{
		Code: constant.SuccessCode,
		Data: data,
		Msg:  msg,
	}
	return result
}

func SuccessData(data interface{}) entity.Result {
	result := entity.Result{
		Code: constant.SuccessCode,
		Data: data,
		Msg:  "success",
	}
	return result
}

func SuccessMsg(msg string) entity.Result {
	result := entity.Result{
		Code: constant.SuccessCode,
		Data: nil,
		Msg:  msg,
	}
	return result
}

func Error(msg string) entity.Result {
	result := entity.Result{
		Code: constant.ErrorCode,
		Data: nil,
		Msg:  msg,
	}
	return result
}
