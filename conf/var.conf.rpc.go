package conf

type RPCConf struct {
	Register string `json:"register" valid:"required"`
}
