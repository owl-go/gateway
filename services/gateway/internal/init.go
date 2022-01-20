package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/grpcclient"
	dis "server/infra/discovery"
	"server/infra/redis"
)

var (
	node        *dis.ServiceNode
	watch       *dis.ServiceWatcher
	redisClient *redis.Redis
)

// Cors 跨域处理
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		//context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,token,X-Requested-With")
		context.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,UPDATE")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
		}
		context.Next()
	}
}

// Init 开启服务
func Init(serviceNode *dis.ServiceNode, ServiceWatcher *dis.ServiceWatcher, cfg *redis.Config) {
	node = serviceNode
	watch = ServiceWatcher
	node.RegisterNode()
	redisClient = redis.NewRedis(cfg)
	go watch.WatchServiceNode("", grpcclient.WatchServiceCallBack)
}
