/*
@author:Deng.l.w
@version:1.20
@date:2023-03-19 11:48
@file:vote.go
*/

package controllers

import (
	"fmt"
	"logic"
	"models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 投票
//
//type VoteData struct {
//	// UserId 请求中获得
//	PostID    int64 `json:"post_id,string"`   //帖子id
//	Direction int   `json:"direction,string"` //赞成票(1) 反对票(-1)
//}

func PostVoteController(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除错误中的结构体
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
	}
	// 获取当前用户的id
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	fmt.Println(p.Direction)
	fmt.Println("errorserrs")
	// 具体投票的业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic vote failed", zap.Error(err))
		ResponseError(c, CodeServerVeryBusy)
		return
	}
	ResponseSuccess(c, nil)
}
