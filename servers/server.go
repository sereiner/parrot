package servers

import (
	"fmt"

	"github.com/sereiner/log"
	"github.com/sereiner/parrot/conf"
)

// ServerType 定义服务器类型
type ServerType byte

const (
	// 目前支持的的服务器

	// APIServer api服务器
	APIServer ServerType = iota
	// RPCServer rpc 服务器
	RPCServer
	// ProxyServer 代理服务器
	ProxyServer
)

//IRegistryServer is based registry server interface,所有类型的服务器都要实现该接口
type IRegistryServer interface {
	Notify(conf.IServerConf) error
	Start() error
	GetAddress() string
	GetServices() []string
	Restarted() bool
	GetStatus() string
	Shutdown()
}

//IServerResolver is server resolvers interface,根据类型产生一个初始化的服务器,所有类型的服务器都要实现该接口
type IServerResolver interface {
	Resolve(registryAddr string, conf conf.IServerConf, log log.ILogger) (IRegistryServer, error)
}

// resolvers 服务器解析容器
var resolvers = make(map[ServerType]IServerResolver)

// Register 根据服务器类型,注册一个服务器
func Register(serverType ServerType, resolver IServerResolver) {

	if _, ok := resolvers[serverType]; ok {
		log.Panicf("服务已经注册过了 %v", serverType)
	}

	resolvers[serverType] = resolver
}

// NewRegistryServer 从服务器解析容器获取一个创建器,创建服务器
func NewRegistryServer(serverType ServerType, registryAddr string, conf conf.IServerConf, log log.ILogger) (IRegistryServer, error) {

	if resolver, ok := resolvers[serverType]; ok {
		return resolver.Resolve(registryAddr, conf, log)
	}

	return nil, fmt.Errorf("创建服务器发生错误,未找到要创建的服务器类型:%v", serverType)
}
