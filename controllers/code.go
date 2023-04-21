/*
@author:Deng.l.w
@version:1.20
@date:2023-03-06 15:16
@file:code.go
*/

package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerVeryBusy
	CodeInvalidToken
	CodeNeedLogin
)

// 自定义错误码对应的错误信息
var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerVeryBusy:  "服务繁忙",
	CodeInvalidToken:    "无效无效的token",
	CodeNeedLogin:       "需要登录",
}

// Msg 返回错误码对应的错误信息
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerVeryBusy]
	}
	return msg
}
