/*
@author:Deng.l.w
@version:1.20
@date:2023-03-07 17:10
@file:post.go
*/

package logic

import (
	"dao/mysql"
	"dao/redis"
	"models"
	"pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//生成postID
	p.ID = int64(snowflake.GenID())

	//入库保存到数据库
	err = mysql.CreatePost(p)
	if err!=nil{
		return err
	}
	
	// 创建时间
	err = redis.CreatePost(p.ID, p.CommunityID)
	if err!=nil{
		return err
	}
	return
	//返回
}

// GetPostById 根据帖子id查询帖子详情数据
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询数据并组合接口想要的数据
	data = new(models.ApiPostDetail)
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid)", zap.Error(err))
		return
	}
	println(post)
	println(post.AuthorID)
	// 根据author_id 查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByid(post.AuthorID)", zap.Error(err))
		return
	}
	//根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error(" mysql.GetCommunityDetailByID(post.CommunityID)", zap.Error(err))
		return
	}
	data.AuthorName = user.Username
	data.CommunityDetail = community
	data.Post = post
	return

}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	data = make([]*models.ApiPostDetail, 0, len(posts))
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByid(post.AuthorID)", zap.Error(err))
			continue
		}
		//根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error(" mysql.GetCommunityDetailByID(post.CommunityID)", zap.Error(err))
			continue
		}
		postdetail := new(models.ApiPostDetail)
		postdetail.AuthorName = user.Username
		postdetail.CommunityDetail = community
		postdetail.Post = post
		data = append(data, postdetail)
	}
	return
}

// GetPostList2 根据发帖时间或分数查询帖子列表
func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}

	// id为0
	if len(ids) == 0 {
		zap.L().Warn("redis GetPostlist")
	}
	//根据id去mysql查询帖子详细信息
	// 返回的数据还要按照给定的id顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	// 提前查好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByid(post.AuthorID)", zap.Error(err))
			continue
		}
		//根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error(" mysql.GetCommunityDetailByID(post.CommunityID)", zap.Error(err))
			continue
		}
		postdetail := new(models.ApiPostDetail)
		postdetail.AuthorName = user.Username
		postdetail.CommunityDetail = community
		postdetail.Post = post
		postdetail.VoteNumber = voteData[idx]
		data = append(data, postdetail)
	}
	return
}

func GetCommunityPostList(p *models.ParamCommunityPostList) (data []*models.ApiPostDetail, err error) {
	// 去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}

	if len(ids) == 0 {
		zap.L().Warn("redis GetPostlist")
	}
	//根据id去mysql查询帖子详细信息
	// 返回的数据还要按照给定的id顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	// 提前查好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByid(post.AuthorID)", zap.Error(err))
			continue
		}
		//根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error(" mysql.GetCommunityDetailByID(post.CommunityID)", zap.Error(err))
			continue
		}
		postdetail := new(models.ApiPostDetail)
		postdetail.AuthorName = user.Username
		postdetail.CommunityDetail = community
		postdetail.Post = post
		postdetail.VoteNumber = voteData[idx]
		data = append(data, postdetail)
	}
	return
}
