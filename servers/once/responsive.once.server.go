package once

import (
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"

	logger "github.com/sereiner/library/log"
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/conf"
	"github.com/sereiner/parrot/engines"
	"github.com/sereiner/parrot/servers"
)

type OnceResponsiveServer struct {
	server        *OnceServer
	engine        servers.IRegistryEngine
	registryAddr  string
	pubs          []string
	shardingIndex int
	shardingCount int
	master        bool
	currentConf   conf.IServerConf
	closeChan     chan struct{}
	once          sync.Once
	done          bool
	pubLock       sync.Mutex
	restarted     bool
	*logger.Logger
	mu sync.Mutex
}

//NewCronResponsiveServer 创建rpc服务器
func NewOnceResponsiveServer(registryAddr string, cnf conf.IServerConf, logger *logger.Logger) (h *OnceResponsiveServer, err error) {
	h = &OnceResponsiveServer{
		closeChan:    make(chan struct{}),
		currentConf:  cnf,
		Logger:       logger,
		pubs:         make([]string, 0, 2),
		registryAddr: registryAddr,
	}
	// 启动执行引擎
	h.engine, err = engines.NewServiceEngine(cnf, registryAddr, h.Logger)
	if err != nil {
		return nil, fmt.Errorf("%s:engine启动失败%v", cnf.GetServerName(), err)
	}
	chandler := cnf.Get("__component_handler_").(component.IComponentHandler)
	if err = h.engine.SetHandler(chandler); err != nil {
		return nil, err
	}
	h.server, err = NewCronServer(h.currentConf.GetServerName(),
		"",
		nil,
		WithShowTrace(cnf.GetBool("trace", false)),
		WithLogger(logger))
	if err != nil {
		return
	}
	err = h.SetConf(true, h.currentConf)
	if err != nil {
		return
	}
	go h.server.Dynamic(h.engine, chandler.GetDynamicCron())
	return
}

//Restart 重启服务器
func (w *OnceResponsiveServer) Restart(cnf conf.IServerConf) (err error) {
	w.Shutdown()
	time.Sleep(time.Second)
	w.closeChan = make(chan struct{})
	w.done = false
	w.currentConf = cnf
	w.once = sync.Once{}
	// 启动执行引擎
	w.engine, err = engines.NewServiceEngine(cnf, w.registryAddr, w.Logger)
	if err != nil {
		return fmt.Errorf("%s:engine启动失败%v", cnf.GetServerName(), err)
	}
	chandler := cnf.Get("__component_handler_").(component.IComponentHandler)
	if err = w.engine.SetHandler(chandler); err != nil {
		return err
	}
	w.server, err = NewCronServer(w.currentConf.GetServerName(),
		"",
		nil,
		WithShowTrace(cnf.GetBool("trace", false)),
		WithLogger(w.Logger))
	if err != nil {
		return
	}
	if err = w.SetConf(true, cnf); err != nil {
		return
	}

	go w.server.Dynamic(w.engine, chandler.GetDynamicCron())

	if err = w.Start(); err == nil {
		w.currentConf = cnf
		w.restarted = true
		return
	}
	return err
}

//Start 启用服务
func (w *OnceResponsiveServer) Start() (err error) {
	err = w.server.Run()
	if err != nil {
		return
	}
	return w.publish()
}

//Shutdown 关闭服务器
func (w *OnceResponsiveServer) Shutdown() {
	w.done = true
	w.once.Do(func() {
		close(w.closeChan)
	})
	w.unpublish()
	w.server.Shutdown(time.Second)
	if w.engine != nil {
		w.engine.Close()
	}
}

//GetAddress 获取服务器地址
func (w *OnceResponsiveServer) GetAddress() string {
	return w.server.GetAddress()
}

//GetStatus 获取当前服务器状态
func (w *OnceResponsiveServer) GetStatus() string {
	return w.server.GetStatus()
}

//GetServices 获取服务列表
func (w *OnceResponsiveServer) GetServices() map[string][]string {
	return w.engine.GetServices()
}

//Restarted 服务器是否已重启
func (w *OnceResponsiveServer) Restarted() bool {
	return w.restarted
}

func (w *OnceResponsiveServer) SetPb(func(component.IContainer, *grpc.Server)) {

}
