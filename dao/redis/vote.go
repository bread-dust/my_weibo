/*
@author:Deng.l.w
@version:1.20
@date:2023-03-19 12:42
@file:vote.go
*/

package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const (
	onWeekInSeconds = 7 * 24 * 3600 // 一周的秒数
	scorePerVote    = 432 // 每票占得分数
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRePested   = errors.New("不允许重复投票")
)

func CreatePost(postID, communityID int64) error {
	pipline := rdb.TxPipeline()
	// 事务 帖子时间
	pipline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipline.ZAdd(getRedisKey(KeyPostScoreZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 把帖子id加到帖子分区
	cKey := getRedisKey(KeyCommunityStPrefix + strconv.Itoa(int(communityID)))
	pipline.SAdd(cKey, postID)
	_, err := pipline.Exec()
	return err

}

func VoteForPost(userID, postID string, direction float64) error {
	// 1. 判断投票限制
	// 去redis 取帖子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	// 判断是否超过一周
	if float64(time.Now().Unix())-postTime > onWeekInSeconds {
		return ErrVoteTimeExpire
	}

	// 2和3放到一个pipline事务
	// 2. 更新分数
	// 查询投票记录
	DirVote := KeyPostVoteZSetPrefix + postID
	oval := rdb.ZScore(getRedisKey(DirVote), userID).Val()
	// oval : -1 0 1
	// 如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if direction == oval {
		return ErrVoteRePested // 重复投票
	}

	var dir float64 // 方向
	if direction > oval {
		dir = 1 
	} else {
		dir = -1 
	}
	diff := math.Abs(oval - direction) // 票的差值
	pipline := rdb.TxPipeline()
	pipline.ZIncrBy(getRedisKey(KeyPostScoreZset), dir * diff * scorePerVote, postID)
	// 为有序集key的成员member的score值加上增量increment

	// 3. 记录用户为该帖子投票的数据
	if direction == 0 {
		pipline.ZRem(getRedisKey(DirVote), userID)
		// 删除有序集key中的一个或多个成员，不存在的成员将被忽略。
	} else {
		pipline.ZAdd(getRedisKey(DirVote), redis.Z{
			Score:  direction, // 用户投的是赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipline.Exec()
	return err
}

// GetCommunityPostListHandler 根据社区查询帖子列表
func GetCommunityPostListHandler(c *gin.Context) {

}
