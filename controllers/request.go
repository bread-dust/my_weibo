/*
@author:Deng.l.w
@version:1.20
@date:2023-03-06 18:58
@file:request.go
*/

package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser 获取当前登录用户的ID
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
