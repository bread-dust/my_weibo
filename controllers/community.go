/*
@author:Deng.l.w
@version:1.20
@date:2023-03-07 15:33
@file:community.go
*/

package controllers

import (
	"logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// 社区相关

func CommunityHandler(c *gin.Context) {
	// 查询所有的社区（community_id,community_name),以切片形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic GetCommunityList failed:err", zap.Error(err))
		ResponseError(c, CodeServerVeryBusy) //不轻易把服务端暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	// 获取社区ID
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64) //10进制，64位
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic GetCommunityList failed:err", zap.Error(err))
		ResponseError(c, CodeServerVeryBusy) //不轻易把服务端暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
