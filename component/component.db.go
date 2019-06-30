package component

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/sereiner/gorose"
	"github.com/sereiner/library/concurrent/cmap"
	"github.com/sereiner/parrot/conf"
	"github.com/sereiner/parrot/registry"
)

//DBTypeNameInVar DB在var配置中的类型名称
const DBTypeNameInVar = "db"

//DBNameInVar DB名称在var配置中的末节点名称
const DBNameInVar = "db"

//IComponentDB Component DB
type IComponentDB interface {
	GetRegularDB(names ...string) (d gorose.IDB)
	GetDB(names ...string) (d gorose.IDB, err error)
	GetDBBy(tpName string, name string) (c gorose.IDB, err error)
	SaveDBObject(tpName string, name string, f func(c conf.IConf) (gorose.IDB, error)) (bool, gorose.IDB, error)
	Close() error
}

//StandardDB db
type StandardDB struct {
	IContainer
	name  string
	dbMap cmap.ConcurrentMap
}

//NewStandardDB 创建DB
func NewStandardDB(c IContainer, name ...string) *StandardDB {
	if len(name) > 0 {
		return &StandardDB{IContainer: c, name: name[0], dbMap: cmap.New(2)}
	}
	return &StandardDB{IContainer: c, name: DBNameInVar, dbMap: cmap.New(2)}
}

//GetRegularDB 获取正式的没有异常数据库实例
func (s *StandardDB) GetRegularDB(names ...string) (d gorose.IDB) {
	d, err := s.GetDB(names...)
	if err != nil {
		panic(err)
	}
	return d
}

//GetDB 获取数据库操作对象
func (s *StandardDB) GetDB(names ...string) (d gorose.IDB, err error) {
	name := s.name
	if len(names) > 0 {
		name = names[0]
	}
	return s.GetDBBy(DBTypeNameInVar, name)
}

//GetDBBy 根据类型获取缓存数据
func (s *StandardDB) GetDBBy(tpName string, name string) (c gorose.IDB, err error) {
	_, c, err = s.SaveDBObject(tpName, name, func(jConf conf.IConf) (gorose.IDB, error) {
		var dbConf conf.DBConf
		if err = jConf.Unmarshal(&dbConf); err != nil {
			return nil, err
		}
		if b, err := govalidator.ValidateStruct(&dbConf); !b {
			return nil, err
		}
		connection, err := gorose.Open(&gorose.DbConfigSingle{
			Driver:          dbConf.Provider, // driver: mysql/sqlite/oracle/mssql/postgres
			EnableQueryLog:  false,           // if enable sql logs
			SetMaxOpenConns: dbConf.MaxOpen,  // connection pool of max Open connections, default zero
			SetMaxIdleConns: dbConf.MaxIdle,  // connection pool of max sleep connections
			Prefix:          "",              // prefix of table
			Dsn:             dbConf.ConnString,
		})
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return connection, err
	})
	return c, err
}

//SaveDBObject 缓存对象
func (s *StandardDB) SaveDBObject(tpName string, name string, f func(c conf.IConf) (gorose.IDB, error)) (bool, gorose.IDB, error) {
	cacheConf, err := s.IContainer.GetVarConf(tpName, name)
	if err != nil {
		return false, nil, fmt.Errorf("%s %v", registry.Join("/", s.GetPlatName(), "var", tpName, name), err)
	}
	key := fmt.Sprintf("%s/%s:%d", tpName, name, cacheConf.GetVersion())
	ok, ch, err := s.dbMap.SetIfAbsentCb(key, func(input ...interface{}) (c interface{}, err error) {
		return f(cacheConf)
	})
	if err != nil {
		err = fmt.Errorf("创建db失败:%s,err:%v", string(cacheConf.GetRaw()), err)
		return ok, nil, err
	}
	return ok, ch.(gorose.IDB), err
}

//Close 释放所有缓存配置
func (s *StandardDB) Close() error {
	s.dbMap.RemoveIterCb(func(k string, v interface{}) bool {
		v.(gorose.IDB).Close()
		return true
	})
	return nil
}
