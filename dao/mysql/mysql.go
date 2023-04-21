/*
@author:Deng.l.w
@version:1.20
@date:2023-02-23 18:39
@file:mysql.go
*/

package mysql

import (
	"settings"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)
import _ "github.com/go-sql-driver/mysql"

var db *sqlx.DB

func Init(cfg *settings.MySOLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("%v\n", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConn) // 最大连接数
	db.SetMaxIdleConns(cfg.MaxIdleConn) // 最大空闲连接数
	return
}

func Close() {
	_ = db.Close()
}
