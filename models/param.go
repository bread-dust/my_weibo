/*
@author:Deng.l.w
@version:1.20
@date:2023-03-05 11:40
@file:param.go
*/

package models

// 帖子排序依据
const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVote 投票参数
type ParamVoteData struct {
	// UserId 请求中获得
	PostID    int64 `json:"post_id,string"`                          //帖子id
	Direction int    `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票(1) 反对票(-1) 取消投票(0)
}

// ParamPostList 帖子列表参数
type ParamPostList struct {
	Page  int64  `form:"page"`
	Size  int64  `form:"size"`
	Order string `form:"order"`
}

// ParamCommunityPostList 按社区获得帖子列表query string 参数
type ParamCommunityPostList struct {
	*ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}
