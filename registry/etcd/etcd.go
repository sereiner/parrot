package etcd

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type EtcdClient struct {
	servers   []string
	timeout   time.Duration
	conn      *clientv3.Client
	eventChan <-chan clientv3.WatchChan
	useCount  int32
	isConnect bool
	once      sync.Once
	closeCh   chan struct{}
	done      bool
}

func New(servers []string, timeout time.Duration) (*EtcdClient, error) {
	client := &EtcdClient{servers: servers, timeout: timeout, useCount: 0}
	client.closeCh = make(chan struct{})
	return client, nil
}

func (e *EtcdClient) Connect() (err error) {
	if e.conn == nil {
		conn, err := clientv3.New(clientv3.Config{
			Endpoints:   e.servers,
			DialTimeout: e.timeout,
		})
		if err != nil {
			return fmt.Errorf("connect failed err:%v ", err)
		}
		e.conn = conn
	}
	atomic.AddInt32(&e.useCount, 1)
	time.Sleep(time.Second)
	e.isConnect = true
	return nil
}

func (e *EtcdClient) IsConnect() bool {
	if e.conn == nil {
		return false
	}
	return e.isConnect
}

func (e *EtcdClient) ReConnect() (err error) {
	e.isConnect = false
	if e.conn != nil {
		e.conn.Close()
	}
	e.done = false
	return e.Connect()
}

func (e *EtcdClient) Close() error {
	atomic.AddInt32(&e.useCount, -1)
	if e.useCount > 0 {
		return nil
	}
	if e.conn != nil {
		e.once.Do(func() {
			e.conn.Close()
		})
	}

	e.isConnect = false
	e.done = true
	e.once.Do(func() {
		close(e.closeCh)
	})
	return nil
}