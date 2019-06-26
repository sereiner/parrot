package rpc

import "testing"

func TestFactoryResolvePath(t *testing.T) {
	def_domain := "NotOnlyBooks_debug"
	def_server := "search_server"
	f := NewInvoker(def_domain, def_server, "zk://127.0.0.1")

	_, d, _, err := f.Request(
		"/search",
		"POST",
		map[string]string{},
		map[string]interface{}{
			"key_word":  []string{"鲁迅"},
			"uid":      "123456",
			"page":      1,
			"page_size": 10,
		},
		true)

	if err != nil {
		t.Log(err)
	}

	t.Log(d)
}
