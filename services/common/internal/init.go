package internal

import (
	"fmt"
	"google.golang.org/grpc"
	dis "server/infra/discovery"
	"server/infra/mysql"
	"server/infra/redis"
	"server/log"
	"sync"
)

var (
	grpcClientConn = make(map[string]*GrpcClientConn)

	lock        sync.RWMutex
	node        *dis.ServiceNode
	watch       *dis.ServiceWatcher
	redisClient *redis.Redis
	mysqlClient *mysql.MysqlDriver
)

// GrpcClientConn grpc连接对象
type GrpcClientConn struct {
	Node *dis.Node
	Conn *grpc.ClientConn
}

func (c *GrpcClientConn) Close() {
	if c.Conn != nil {
		c.Conn.Close()
		c.Conn = nil
	}
}

func NewGrpcConnection(node dis.Node) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", node.Nip, node.Port), grpc.WithInsecure())
	if err != nil {
		return nil
	}
	return conn
}

// GetGrpcClientConn 获取一个grpc连接
func GetGrpcClientConn(version string, name string) *grpc.ClientConn {
	lock.RLock()
	defer lock.RUnlock()
	for _, c := range grpcClientConn {
		if c.Node.Version == version && c.Node.Name == name {
			return c.Conn
		}
	}
	return nil
}

// 释放所有连接
func DelAllGrpcClientConn() {
	lock.RLock()
	defer lock.RUnlock()
	for _, c := range grpcClientConn {
		c.Close()
	}
	grpcClientConn = make(map[string]*GrpcClientConn)
}

// Free 关闭服务
func Free() {
	if node != nil {
		node.Close()
		node = nil
	}
	if watch != nil {
		watch.Close()
		watch = nil
	}
	DelAllGrpcClientConn()
}

// Init 初始化服务
func Init(serviceNode *dis.ServiceNode, ServiceWatcher *dis.ServiceWatcher, cfg *redis.Config, mCfg *mysql.MysqlConfig) {
	node = serviceNode
	watch = ServiceWatcher
	redisClient = redis.NewRedis(cfg)
	mysqlClient = mysql.NewMysqlDriver(mCfg)

	go watch.WatchServiceNode("", WatchServiceCallBack)

}

// WatchServiceCallBack 查看所有的Node节点
func WatchServiceCallBack(state dis.NodeStateType, node dis.Node) {
	lock.Lock()
	defer lock.Unlock()
	if state == dis.ServerUp {
		log.Infof("WatchServiceCallBack node up %v", node)
		nid := node.Nid
		_, found := grpcClientConn[nid]
		if !found {
			nodeConn := &GrpcClientConn{
				Node: &node,
				Conn: NewGrpcConnection(node),
			}
			grpcClientConn[nid] = nodeConn
		}
	} else if state == dis.ServerDown {
		log.Infof("WatchServiceCallBack node down %v", node.Nid)
		nodeConn, found := grpcClientConn[node.Nid]
		if found {
			nodeConn.Close()
			delete(grpcClientConn, node.Nid)
		}
	}
}
