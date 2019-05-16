package zookeeper

import (
	"fmt"
	"time"

	"github.com/sereiner/parrot/registry"
	logger "github.com/sereiner/log"
	"github.com/sereiner/lib/zk"
)

//zkRegistry 基于zookeeper的注册中心
type zkRegistry struct {
}

//Resolve 根据配置生成zookeeper客户端
func (z *zkRegistry) Resolve(servers []string, u string, p string, log *logger.Logger) (registry.IRegistry, error) {
	if len(servers) == 0 {
		return nil, fmt.Errorf("未指定zk服务器地址")
	}
	zclient, err := zk.NewWithLogger(servers, time.Second, log, zk.WithdDigest(u, p))
	if err != nil {
		return nil, err
	}
	err = zclient.Connect()
	return zclient, err
}

func init() {
	registry.Register("zk", &zkRegistry{})
}
