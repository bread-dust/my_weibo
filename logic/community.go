/*
@author:Deng.l.w
@version:1.20
@date:2023-03-07 15:39
@file:community.go
*/

package logic

import (
	"dao/mysql"
	"models"
)

// GetCommunityList 查询所有社区id，name
func GetCommunityList() (communityList []*models.Community, err error) {
	//查数据库，找到所有的community
	return mysql.GetCommunityList()
}

// GetCommunityDetail 根据社区ID查询社区详情
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
