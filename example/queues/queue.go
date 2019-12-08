package queues

import (
	"encoding/json"
	"fmt"
	"github.com/sereiner/library/types"
	"github.com/sereiner/parrot/component"
)

//Send 发送消息信息
func Send(c component.IContainer, name string, content string) error {
	q := c.GetRegularQueue()
	return q.Push(name, content)
}

//SendList 发送多组mq队列
func SendList(c component.IContainer, params ...interface{}) (err error) {
	//检查数据
	if len(params)%2 != 0 {
		return fmt.Errorf("参数须成对")
	}

	//发送mq信息
	q := c.GetRegularQueue()
	for k, v := range params {
		if k%2 == 0 {
			tp, er := json.Marshal(params[k+1].(map[string]interface{}))
			if er != nil {
				err = er
				continue
			}

			//fmt.Println("string(tp):", string(tp))
			er = q.Push(types.GetString(v), string(tp))
			if er != nil {
				err = er
			}
		}
	}

	return err
}
