/*
@author:Deng.l.w
@version:1.20
@date:2023-03-19 11:35
@file:keys.go.go
*/

package redis

// redis key 
// redis key注意使用命名空间，方便业务拆分

const (
	KeyPrefix             = "bluebell:"   // 项目key前缀
	KeyPostTimeZSet       = "post:time"   // zset;帖子及发帖时间
	KeyPostScoreZset       = "post:score"  // zset 帖子及投票的分数
	KeyPostVoteZSetPrefix = "post:voted:" //zset记录用户及投票类型；参数是post_id
	KeyCommunityStPrefix  = "community:"  //set;保存每个分区下的帖子
)

// 给redis加前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
