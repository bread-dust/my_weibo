/*
@author:Deng.l.w
@version:1.20
@date:2023-03-06 15:10
@file:response.go
*/

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	{
		"code":10001 //程序的错误码
		“msg":xx, //提示信息
		"data":{} // 存放数据
*/

type ResponseData struct {
	Code ResCode `json:"code"`
	Msg  any     `json:"msg"`
	Data any     `json:"data,omitempty"`
}

// ResponseError 不带信息的响应错误
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 自定义带信息的响应错误
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 自定义响应成功信息
func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
