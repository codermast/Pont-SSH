package utils

import "PontSsh/backend/constant"

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Success(data interface{}, msg string) Result {
	result := Result{
		Code: constant.SuccessCode,
		Data: data,
		Msg:  msg,
	}
	return result
}

func SuccessData(data interface{}) Result {
	result := Result{
		Code: constant.SuccessCode,
		Data: data,
		Msg:  "success",
	}
	return result
}

func SuccessMsg(msg string) Result {
	result := Result{
		Code: constant.SuccessCode,
		Data: nil,
		Msg:  msg,
	}
	return result
}

func Error(msg string) Result {
	result := Result{
		Code: constant.ErrorCode,
		Data: nil,
		Msg:  msg,
	}
	return result
}
