/*
@author:Deng.l.w
@version:1.20
@date:2023-03-28 20:56
@file:post.go
*/

package redis

import (
	"models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func GetIDsFromKey(key string, page, size int64) ([]string, error) {
	// 确定查询的索引起始点
	start := (page - 1) * size
	end := start + size - 1

	// ZREVRANGE 查询 按分数从大到小的顺序查询指定数量的元素
	return rdb.ZRevRange(key, start, end).Result()
}


func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 根据用户请求中携带的order参数
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZset)
	}

	// 确定查询的索引起始点
	return GetIDsFromKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVoteZSetPrefix + id)
	//	// 查找key中分数是1 的数量
	//	v1 := rdb.ZCount(key, "1", "1").Val()
	//	data = append(data, v1)
	//}

	//keys := make([]string, 0, len(ids))
	pipline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVoteZSetPrefix + id)
		pipline.ZCount(key, "1", "1")
	}
	cmders, err := pipline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, err
}

// GetCommunityPostIDsInOrder 按社区查询 ids
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	// 使用zinterstore 把分区的帖子set与帖子分数的zset生成一个新的zset
	// 针对新的zset按之前的逻辑取数据
	// 利用缓存key减少zintersore 执行的次数
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZset)
	}

	key := orderKey + strconv.Itoa(int(p.CommunityID))

	// 社区的key
	cKey := getRedisKey(KeyCommunityStPrefix + strconv.Itoa(int(p.CommunityID)))
	if rdb.Exists(key).Val() < 1 {
		// 不存在，需要计算
		pipline := rdb.Pipeline()
		pipline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey) // zinterstore 计算
		pipline.Expire(key, 60*time.Second) // 超时时间
		_, err := pipline.Exec()
		if err != nil {
			return nil, err
		}
	}

	// 存在的话根据key查ids
	return GetIDsFromKey(key, p.Page, p.Size)

}
