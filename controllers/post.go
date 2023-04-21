/*
@author:Deng.l.w
@version:1.20
@date:2023-03-07 16:46
@file:post.go
*/

package controllers

import (
	"logic"
	"models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 获取参数
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("shouldbind err", zap.Any("err", err))
		zap.L().Error("create post with invalidparam", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 从c中取到当前发请求的用户ID
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	//创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("login createPost failed ", zap.Error(err))
		ResponseError(c, CodeServerVeryBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获得帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	// 1.获取参数（从url帖子id,postid)
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get postdetail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 根据id取出帖子数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic getpostid failed", zap.Error(err))
		return
	}
	// 3， 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获得帖子详情列表的处理函数
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	pageNumStr := c.Query("page")
	pageSizeStr := c.Query("size")
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		page = 1
	}

	size, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		size = 0
	}
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList()", zap.Error(err))
		ResponseError(c, CodeServerVeryBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版，根据前端传来的参数（分数、创建时间）动态排序
// 获取参数
// 去redis 查询id列表
// 根据id去数据库查询帖子详细信息
func GetPostListHandler2(c *gin.Context) {
	// 初始化结构体指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	// 获取分页参数
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取数据
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList()", zap.Error(err))
		ResponseError(c, CodeServerVeryBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// GetCommunityPostListHandler 根据社区查询帖子列表
func GetCommunityPostListHandler(c *gin.Context) {
	// 初始化结构体指定初始参数
	p := &models.ParamCommunityPostList{
		ParamPostList: &models.ParamPostList{
			Page:  1,
			Size:  10,
			Order: models.OrderTime,
		},
	}
	// 获取分页参数
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取数据
	data, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("logic.GetPostList()", zap.Error(err))
		ResponseError(c, CodeServerVeryBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}
