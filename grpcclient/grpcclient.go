package grpcclient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	dis "server/infra/discovery"
	"server/log"
	"server/proto"
	"sync"
	"time"
)

var (
	grpcClientConn = make(map[string]*GrpcClientConn)
	lock           sync.RWMutex
)

// GrpcClientConn grpc连接对象
type GrpcClientConn struct {
	Node *dis.Node
	Conn *grpc.ClientConn
}

// Close 关闭grpc连接
func (c *GrpcClientConn) Close() {
	if c.Conn != nil {
		c.Conn.Close()
		c.Conn = nil
	}
}

// NewGrpcConnection 新建一个grpc连接
func NewGrpcConnection(node *dis.Node) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", node.Nip, node.Port), grpc.WithInsecure())
	if err != nil {
		return nil
	}
	return conn
}

// GetGrpcClientConn 获取一个grpc连接
func GetGrpcClientConn(version string, name string) (*grpc.ClientConn, bool) {
	lock.RLock()
	defer lock.RUnlock()
	for _, c := range grpcClientConn {
		if c.Node.Version == version && c.Node.Name == name {
			return c.Conn, true
		}
	}
	return nil, false
}

func GetGrpcClientByNid(nid string) (*GrpcClientConn, bool) {
	lock.RLock()
	defer lock.RUnlock()
	grpcClinet, ok := grpcClientConn[nid]
	return grpcClinet, ok
}

func AddGrpcClient(node *dis.Node) {
	lock.Lock()
	defer lock.Unlock()
	nodeConn := &GrpcClientConn{
		Node: node,
		Conn: NewGrpcConnection(node),
	}
	grpcClientConn[node.Nid] = nodeConn
}

func DeleteGrpcClinet(nid string) {
	lock.Lock()
	defer lock.Unlock()
	delete(grpcClientConn, nid)
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
	DelAllGrpcClientConn()
}

// WatchServiceCallBack 查看所有的Node节点
func WatchServiceCallBack(state dis.NodeStateType, node dis.Node) {
	if state == dis.ServerUp {
		log.Infof("WatchServiceCallBack node up %v", node.Nid)
		nid := node.Nid
		_, found := GetGrpcClientByNid(nid)
		if !found {
			AddGrpcClient(&node)
		}
	} else if state == dis.ServerDown {
		log.Infof("WatchServiceCallBack node down %v", node.Nid)
		nodeConn, found := GetGrpcClientByNid(node.Nid)
		if found {
			nodeConn.Close()
			DeleteGrpcClinet(node.Nid)
		}
	}
}

func findGrpcClientConnByName(serviceName string) (*grpc.ClientConn, bool) {
	lock.RLock()
	defer lock.RUnlock()
	for _, c := range grpcClientConn {
		if c.Node.Name == serviceName {
			return c.Conn, true
		}
	}
	return nil, false
}

func CallService(version, serviceName string, req *proto.Request) (*proto.Response, error, bool) {
	var conn *grpc.ClientConn
	var ok bool
	if version == "" || version == "*" {
		conn, ok = findGrpcClientConnByName(serviceName)
		if !ok {
			return nil, nil, false
		}
	} else {
		conn, ok = GetGrpcClientConn(version, serviceName)
		if !ok {
			return nil, nil, false
		}
	}
	if conn == nil {
		return nil, nil, false
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		client := proto.NewCallServiceClient(conn)
		r, err := client.CallService(ctx, req)
		return r, err, true
	}
}
