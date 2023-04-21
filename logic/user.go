/*
@author:Deng.l.w
@version:1.20
@date:2023-03-05 11:22
@file:user.go
*/

package logic

import (
	"dao/mysql"
	"models"
	"pkg/jwt"
	"pkg/snowflake"
)

// SignUp 存放注册业务逻辑代码
func SignUp(mp *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err = mysql.CheckUserExist(mp.Username); err != nil {
		//查询出错
		return err
	}

	// 1.生成UID

	userID := snowflake.GenID()
	// 构造一个User
	u := models.User{
		UserID: userID,
		Username: mp.Username,
		Password: mp.Password,
	}
	// 2. 密码加密
	// 3.保存进数据库
	return mysql.InsertUser(&u)
}

// Login 存放登录业务逻辑代码
func Login(mp *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: mp.Username,
		Password: mp.Password,
	}

	// 传递的是指针，能拿到userid
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	//生成jwt
	return jwt.GenToken(user.UserID, user.Username)
}
