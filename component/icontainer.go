package component

import (
	"github.com/sereiner/gorose"
	"github.com/sereiner/library/cache"
	"github.com/sereiner/library/influxdb"
	"github.com/sereiner/library/queue"
	logger "github.com/sereiner/log"
	"github.com/sereiner/parrot/conf"
	"github.com/sereiner/parrot/context"
	"github.com/sereiner/parrot/registry"
)

type IContainer interface {
	context.RPCInvoker
	GetComponent() IComponent
	conf.ISystemConf
	conf.IVarConf
	conf.IMainConf
	GetRegistry() registry.IRegistry
	GetLogger() logger.ILogging

	GetRegularCache(names ...string) (c cache.ICache)
	GetCache(names ...string) (c cache.ICache, err error)
	GetCacheBy(tpName string, name string) (c cache.ICache, err error)
	SaveCacheObject(tpName string, name string, f func(c conf.IConf) (cache.ICache, error)) (bool, cache.ICache, error)

	GetRegularDB(names ...string) (d gorose.IDB)
	GetDB(names ...string) (d gorose.IDB, err error)
	GetDBBy(tpName string, name string) (c gorose.IDB, err error)
	SaveDBObject(tpName string, name string, f func(c conf.IConf) (gorose.IDB, error)) (bool, gorose.IDB, error)

	GetRegularInflux(names ...string) (c influxdb.IInfluxClient)
	GetInflux(names ...string) (d influxdb.IInfluxClient, err error)
	GetInfluxBy(tpName string, name string) (c influxdb.IInfluxClient, err error)
	SaveInfluxObject(tpName string, name string, f func(c conf.IConf) (influxdb.IInfluxClient, error)) (bool, influxdb.IInfluxClient, error)

	GetRegularQueue(names ...string) (c queue.IQueue)
	GetQueue(names ...string) (q queue.IQueue, err error)
	GetQueueBy(tpName string, name string) (c queue.IQueue, err error)
	SaveQueueObject(tpName string, name string, f func(c conf.IConf) (queue.IQueue, error)) (bool, queue.IQueue, error)

	GetGlobalObject(tpName string, name string) (c interface{}, err error)
	SaveGlobalObject(tpName string, name string, f func(c conf.IConf) (interface{}, error)) (bool, interface{}, error)
	Close() error
}
