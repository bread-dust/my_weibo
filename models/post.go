/*
@author:Deng.l.w
@version:1.20
@date:2023-03-07 16:47
@file:post.go
*/

package models

import "time"

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id,string" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
}

// ApiPostDetail 帖子详情接口
type ApiPostDetail struct {
	AuthorName       string                    `json:"author_name"`
	VoteNumber       int64                     `json:"votes"`
	*Post            `json:"post"`             //帖子结构体
	*CommunityDetail `json:"community_detail"` // 社区信息结构体
}
