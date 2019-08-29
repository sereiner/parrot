package ws

import (
	logger "github.com/sereiner/library/log"
	"github.com/sereiner/parrot/conf"
	"github.com/sereiner/parrot/servers"
)

type wsServerAdapter struct {
}

func (h *wsServerAdapter) Resolve(registryAddr string, conf conf.IServerConf, log *logger.Logger) (servers.IRegistryServer, error) {
	return NewWSServerResponsiveServer(registryAddr, conf, log)
}

func init() {
	servers.Register("ws", &wsServerAdapter{})
}
