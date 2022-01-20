package discovery

import (
	"encoding/json"
	"fmt"
)

// NodeStateType 节点状态
type NodeStateType int32

const (
	// ServerUp 服务存活
	ServerUp NodeStateType = 0
	// ServerDown 服务死亡
	ServerDown NodeStateType = 1
)

// Node 服务节点对象
type Node struct {
	Version string //版本
	Region  string //地区
	Zone    string //分区
	Name    string //服务名称
	Nid     string //节点id
	Nip     string //节点ip
	Port    string //节点端口
	Payload int    //负载
}

// GetNodeValue 获取节点保存的值
func (node *Node) GetNodeValue() string {
	info, err := json.Marshal(node)
	if err != nil {
		return ""
	}
	return string(info)
}

// Encode 将map格式转换成string
func Encode(data map[string]string) string {
	if data != nil {
		str, _ := json.Marshal(data)
		return string(str)
	}
	return ""
}

// Decode 将string格式转换成map
func Decode(str []byte) map[string]string {
	if len(str) > 0 {
		var data map[string]string
		json.Unmarshal(str, &data)
		return data
	}
	return nil
}

// GetEventChannel 获取广播对象string
func GetEventChannel(node Node) string {
	return "event-" + node.Nid
}

// GetRPCChannel 获取RPC对象string
func GetRPCChannel(node Node) string {
	return "rpc-" + node.Nid
}

func generateKey(node Node) string {
	return fmt.Sprintf("/%s/%s/%s/%s/%s", node.Version, node.Region, node.Zone, node.Name, node.Nid)
}
