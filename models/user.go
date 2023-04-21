/*
@author:Deng.l.w
@version:1.20
@date:2023-03-05 17:18
@file:user.go
*/

package models

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
