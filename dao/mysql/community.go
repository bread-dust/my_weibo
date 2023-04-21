/*
@author:Deng.l.w
@version:1.20
@date:2023-03-07 15:42
@file:community.go
*/

package mysql

import (
	"models"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

var (
	ErrUserExist       = errors.New("用户已存在")
	ErrUserNotExist    = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("用户名或密码错误")
	ErrInvalidID       = errors.New("无效的ID")
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := `select 
    community_id,community_name,introduction,create_time 
from community 
where community_id = ?
`
	if err = db.Get(communityDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrInvalidID
		}
	}
	return communityDetail, err
}
