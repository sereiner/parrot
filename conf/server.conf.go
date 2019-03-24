package conf

//IServerConf server conf interface
type IServerConf interface {
	ISystemConf
}

type ISystemConf interface {
	GetPlatName() string
	GetSysName() string
	GetServerType() string
	GetClusterName() string
	GetServerName() string
	GetAppConf(v interface{}) error
	Get(key string) interface{}
	Set(key string, value interface{})
}
