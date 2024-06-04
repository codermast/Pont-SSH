package utils

import "PontSsh/backend/constant"

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Success(data interface{}, msg string) Result {
	Result := Result{
		Code: constant.SuccessCode,
		Data: data,
		Msg:  msg,
	}
	return Result
}

func SuccessData(data interface{}) Result {
	Result := Result{
		Code: constant.SuccessCode,
		Data: data,
		Msg:  "success",
	}
	return Result
}

func SuccessMsg(msg string) Result {
	Result := Result{
		Code: constant.SuccessCode,
		Data: nil,
		Msg:  msg,
	}
	return Result
}
