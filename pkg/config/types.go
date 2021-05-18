package config

import "net"

type Cfg interface {
	GetNodeByMACAddr(macAddr string) (*Node, error)
}

type NodeType int32

const (
	TypeWorkerNode NodeType = 0
	TypeMasterNode NodeType = 1
)

func IsWorkerNode(nt NodeType) bool {
	if nt == TypeWorkerNode {
		return true
	}
	return false
}

func IsMasterNode(nt NodeType) bool {
	if nt == TypeMasterNode {
		return true
	}
	return false
}

var (
	NodeTypeName = map[int32]string{
		0: "worker",
		1: "master",
	}
	NodeTypeValue = map[string]int32{
		"worker": 0,
		"master": 1,
	}
)

type ClusterNetwork struct {
	Gateway string   `json:"gateway"`
	Netmask string   `json:"netmask"`
	DNS     []string `json:"dns"`
}

type Cluster struct {
	Name    string         `json:"name"`
	Network ClusterNetwork `json:"network"`
	Workers []Node         `json:"workers"`
	Masters []Node         `json:"masters"`
}

func (cluster *Cluster) AllNodes() []Node {
	allNodes := append(cluster.Masters, cluster.Workers...)
	return allNodes
}

type Node struct {
	Name    string `json:"name"`
	MACAddr string `json:"macAddr"`
	IPAddr  string `json:"ipAddr"`
}

func (node *Node) GetName() string {
	return node.Name
}

func (node *Node) GetMACAddr() net.HardwareAddr {
	mac, _ := net.ParseMAC(node.MACAddr)
	return mac
}

func (cluster *Cluster) GetNode(nodeName string) (*Node, error) {
	for _, node := range cluster.AllNodes() {
		if node.GetName() == nodeName {
			return &node, nil
		}
	}
	return nil, ErrNodeNotFound
}

func (cluster *Cluster) GetNodeByMACAddr(macAddr string) (*Node, error) {
	_, err := net.ParseMAC(macAddr)
	if err != nil {
		return nil, err
	}
	for _, node := range cluster.AllNodes() {
		if node.MACAddr == macAddr {
			return &node, nil
		}
	}
	return nil, ErrNodeNotFound

}
