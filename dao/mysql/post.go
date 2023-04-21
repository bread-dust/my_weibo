/*
@author:Deng.l.w
@version:1.20
@date:2023-03-07 17:12
@file:post.go
*/

package mysql

import (
	"models"
	"strings"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id,title,content,author_id,community_id)
				values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err!=nil{
		zap.L().Error("error",zap.Error(err))
	}
	return 
}

// GetPostById 根据id查询单个帖子的详情数据
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
				from post 
				where post_id=?`
	err = db.Get(post, sqlStr, pid)
	return 
}

// GetPostList 查询帖子列表
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
				from post 
				order by create_time
				desc
				limit ?,?`

	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// GetPostListByIDs 根据id列表查询帖子列表
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	// sqlx.in 查询集合内的数据
	// Find_In_Set() 给定顺序排序 
	sqlStr := `select psot_id,title,content,author_id,community_id,create_time 
				from post
				where post_id in (?)
				order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query) // ！！！ 需要重新绑定

	db.Select(&postList, query, args...) // ！！！ 需要三个引号
	return
}
