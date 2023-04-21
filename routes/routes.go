/*
@author:Deng.l.w
@version:1.20
@date:2023-03-04 0:12
@file:routes.go
*/

package routes

import (
	"controllers"
	"logger"
	"midware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true),midware.RatelimitMiddleware(2*time.Second,4))

	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	// 登录业务路由
	v1.POST("/login", controllers.LoginHandler)

	//应用jwt认证中间件
	v1.Use(midware.JWTAuthMiddleware())

	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts", controllers.GetPostListHandler)

		// 根据时间或分数获取帖子列表
		v1.GET("/posts2", controllers.GetPostListHandler2)
		// 投票
		v1.POST("/vote", controllers.PostVoteController)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})
	return r
}
