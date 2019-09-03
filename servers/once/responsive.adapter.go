package once

import (
	logger "github.com/sereiner/library/log"
	"github.com/sereiner/parrot/conf"
	"github.com/sereiner/parrot/servers"
)

type rpcServerAdapter struct {
}

func (h *rpcServerAdapter) Resolve(c string, conf conf.IServerConf, log *logger.Logger) (servers.IRegistryServer, error) {
	return NewOnceResponsiveServer(c, conf, log)
}

func init() {
	servers.Register("once", &rpcServerAdapter{})
}
