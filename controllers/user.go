/*
@author:Deng.l.w
@version:1.20
@date:2023-03-05 11:18
@file:user.go
*/

package controllers

import (
	msq "dao/mysql"
	"logic"
	"models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求参数
func SignUpHandler(c *gin.Context) {
	// 1.参数处理
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数出错
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		//判断err是不是validator.ValidationErrors类型
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 手动对请求参数进行详细的业务规则判断
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	fmt.Println(&p)
	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, msq.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		}
		ResponseError(c, CodeServerVeryBusy)
		return

	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录逻辑处理
func LoginHandler(c *gin.Context) {
	// 获取参数及校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数出错
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		//判断err是不是validator.ValidationErrors类型
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, msq.ErrorInvalidPassword) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 返回响应
	ResponseSuccess(c, token)
}
