package rpc

import "testing"



func TestRPCClient(t *testing.T) {
	var invoker = NewInvoker("NotOnlyBooks_debug", "search_server", "192.168.0.105:9002")

	invoker.PreInit("/test")
	_, d, _, err  := invoker.Request(
		"/test",
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
	//b.ResetTimer()
	//for i := 0; i < b.N; i++ {
	//	s, _, _, _ := invoker.Request(
	//		"/order/request/success",
	//		make(map[string]string),
	//		true)
	//
	//}
}
