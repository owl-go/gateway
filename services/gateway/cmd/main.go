package main

import (
	"fmt"
	"net/http"
	"server/grpcclient"
	dis "server/infra/discovery"
	"server/infra/redis"
	"server/services/gateway/internal"
	"server/utils"
	"strconv"

	lg "server/log"
	"server/services/gateway/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	defer stop()
	lg.Init(conf.Log.Level)
	go debug()
	start()
	httpRouter()
}

func start() {
	serviceNode := dis.NewServiceNode(utils.ProcessUrlString(conf.Etcd.Addrs),
		conf.Global.Version, conf.Global.Region, conf.Global.Zone,
		conf.Global.Name, conf.Global.Nid, conf.Global.Nip, strconv.Itoa(conf.Global.Port))
	serviceWatcher := dis.NewServiceWatcher(utils.ProcessUrlString(conf.Etcd.Addrs))

	cfg := redis.Config{
		Addrs: conf.Redis.Addrs,
		Pwd:   conf.Redis.Password,
		DB:    conf.Redis.DB,
	}

	internal.Init(serviceNode, serviceWatcher, &cfg)
}

func stop() {
	grpcclient.Free()
}

func httpRouter() {
	router := gin.Default()
	router.Use(internal.Cors())

	router.POST("/:v/:service/:method", internal.PostHandler)
	router.POST("/:v/:service/:method/:submethod", internal.PostHandler)
	router.POST("/:v/:service/:method/:submethod/:extmethod", internal.PostHandler)
	//router.PUT("/:v/:service/:method", internal.PutHandler)
	router.GET("/:v/:service/:method", internal.GetHandler)
	router.GET("/:v/:service/:method/:submethod", internal.GetHandler)
	router.GET("/:v/:service/:method/:submethod/:extmethod", internal.GetHandler)
	//router.DELETE("/:v/:service/:method", internal.DeleteHandler)
	router.Run(fmt.Sprintf("%s:%d", conf.Global.Nip, conf.Global.Port))
}

func debug() {
	if conf.Global.Pprof != "" {
		lg.Infof("start gateway pprof on %s", conf.Global.Pprof)
		http.ListenAndServe(conf.Global.Pprof, nil)
	}
}
