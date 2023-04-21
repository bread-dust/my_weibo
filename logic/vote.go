/*
@author:Deng.l.w
@version:1.20
@date:2023-03-19 11:37
@file:vote.go
*/

package logic

import (
	"dao/redis"
	"fmt"
	"models"
	"strconv"

	"go.uber.org/zap"
)

// 投票功能：
// 1. 用户投票的数据

// 投一票加432分,86400/200 -> 需要200 张赞成票可以给帖子续一天

/*
投票的几种情况：
direction=1,两种情况
	1.之前没有投过票，现在投赞成票 -> 更新分数和记录投票 差值绝对值：1 +432
	2. 之前投反对票，现在改赞成票-> 更新分数和记录投票 差值绝对值：2 +432*2
direction=0,两种情况
	1.之前投过赞成票，现在要取消投票-> 更新分数和记录投票 差值绝对值：1 -432
	2.之前投过反对票，现在要取消投票-> 更新分数和记录投票 差值绝对值：1 +432
durection=-1，两种情况
	1.之前没有投过票，现在投反对票-> 更新分数和记录投票 差值绝对值：1 -432
	2.之前投过赞成票，现在投反对票-> 更新分数和记录投票 差值绝对值：2 -432*2

投票限制：
	每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期不允许投票
	1.到期之后将redis中保存的赞成票数及反对票数存储到mysql中
	2.到期之后删除 KeyPostVotedZSet
*/

// VoteForPost 为帖子投票的函数
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	postIDString:= strconv.Itoa(int(p.PostID))
	userIDString:= strconv.Itoa(int(userID))
	zap.L().Debug("voteforPost",
		zap.String("userid", userIDString),
		zap.String("postid", postIDString),
		zap.Int("direction", p.Direction))
	fmt.Println(p.Direction)
	return redis.VoteForPost(userIDString, postIDString, float64(p.Direction))

}
