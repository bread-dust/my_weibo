/*
@author:Deng.l.w
@version:1.20
@date:2023-03-05 10:06
@file:snowflake.go
*/

package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

// 初始化全局变量Node节点
var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time // 时间因子
	st, err = time.Parse("2006-01-02", startTime)
	// 时间因子从那一年开始
	if err != nil {
		return
	}
	// 项目开始的时间戳，毫秒
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID) // 根据机器ID生成雪花Node节点
	return err
}

func GenID() int64 {
	return node.Generate().Int64() // 生成雪花ID的64位整数
}
