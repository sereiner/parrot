package cron

import (
	"github.com/sereiner/parrot/conf"
	"github.com/sereiner/parrot/servers"
	logger "github.com/sereiner/log"
)

type rpcServerAdapter struct {
}

func (h *rpcServerAdapter) Resolve(c string, conf conf.IServerConf, log *logger.Logger) (servers.IRegistryServer, error) {
	return NewCronResponsiveServer(c, conf, log)
}

func init() {
	servers.Register("cron", &rpcServerAdapter{})
}
