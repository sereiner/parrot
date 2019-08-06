package ws

import (
	"github.com/sereiner/parrot/conf"
	"github.com/sereiner/parrot/servers"
	logger "github.com/sereiner/library/log"
)

type wsServerAdapter struct {
}

func (h *wsServerAdapter) Resolve(registryAddr string, conf conf.IServerConf, log *logger.Logger) (servers.IRegistryServer, error) {
	return NewWSServerResponsiveServer(registryAddr, conf, log)
}

func init() {
	servers.Register("ws", &wsServerAdapter{})
}
