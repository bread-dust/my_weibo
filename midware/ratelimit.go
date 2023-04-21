package midware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/juju/ratelimit"
)

// RatelimitMiddleware 令牌桶限流中间件
func RatelimitMiddleware(fillInterval time.Duration,cap int64) func(c *gin.Context){
	bucket := ratelimit.NewBucket(fillInterval,cap)
	return func(c *gin.Context){
		// 如果桶中没有token，则返回429
		if bucket.TakeAvailable(1) > 0{
			c.String(http.StatusOK,"rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}