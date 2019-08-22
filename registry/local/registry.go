package local

import (
	"github.com/sereiner/parrot/registry"
	logger "github.com/sereiner/library/log"
)

//zkRegistry 基于zookeeper的注册中心
type lkRegistry struct {
}

//Resolve 根据配置生成zookeeper客户端
func (z *lkRegistry) Resolve(servers []string, u string, p string, log logger.ILogging) (registry.IRegistry, error) {
	prefix := "."
	if len(servers) > 0 {
		prefix = servers[0]
	}
	client, err := newLocal(prefix)
	if err != nil {
		return nil, err
	}
	client.Start()
	return client, nil
}

func init() {
	registry.Register("fs", &lkRegistry{})
}
