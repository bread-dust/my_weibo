/*
@author:Deng.l.w
@version:1.20
@date:2023-03-06 18:47
@file:auth.go
*/

package midware

import (
	"controllers"
	"pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	//客户端携带token的三种方式：请求头，请求体，url
	//此处假设token放在header，bearer开头
	// Authorization:bearer xxxx.xxxx.xx X-TOKEN:xxx.xxx.xx
	//具体实现根据业务
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格切割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// part1 是token值，使用jwt解析函数解析
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		//将当前请求的，保存到请求的userID上下文中
		c.Set(controllers.CtxUserKey, mc.UserID)
		c.Next() //后续处理函数通过c.Get获取当前的清气去信息
	}
}
