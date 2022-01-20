package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"server/infra/mysql"
	"server/infra/redis"
	"strconv"

	dis "server/infra/discovery"
	lg "server/log"
	"server/proto"
	"server/services/common/conf"
	"server/services/common/internal"
	"server/utils"

	"google.golang.org/grpc"
)

func start() {
	lg.Init(conf.Log.Level)
	if conf.Global.Pprof != "" {
		go func() {
			lg.Infof("Start pprof on %s", conf.Global.Pprof)
			http.ListenAndServe(conf.Global.Pprof, nil)
		}()
	}
	serviceNode := dis.NewServiceNode(utils.ProcessUrlString(conf.Etcd.Addrs), conf.Global.Version, conf.Global.Region,
		conf.Global.Zone, conf.Global.Name, conf.Global.Nid, conf.Global.Nip, strconv.Itoa(conf.Global.Port))
	serviceNode.RegisterNode()
	serviceWatcher := dis.NewServiceWatcher(utils.ProcessUrlString(conf.Etcd.Addrs))

	rcfg := redis.Config{
		Addrs: conf.Redis.Addrs,
		Pwd:   conf.Redis.Password,
		DB:    conf.Redis.DB,
	}
	mcfg := mysql.MysqlConfig{
		Host:     conf.Mysql.Host,
		Port:     conf.Mysql.Port,
		Username: conf.Mysql.Username,
		Password: conf.Mysql.Password,
		Database: conf.Mysql.DB,
	}
	internal.Init(serviceNode, serviceWatcher, &rcfg, &mcfg)
}

func main() {
	start()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.Global.Nip, conf.Global.Port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	server := new(internal.Server)
	server.Init()

	proto.RegisterCallServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if er := s.Serve(lis); err != nil {
		log.Fatal(er)
	}
}
