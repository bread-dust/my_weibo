/*
@author:Deng.l.w
@version:1.20
@date:2023-03-07 15:50
@file:community.go
*/

package models

import "time"

type Community struct {
	ID   int64  `json:"id,string" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id,string" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
