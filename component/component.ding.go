package component

import (
	"fmt"
	"github.com/sereiner/library/net/http"
	"github.com/sereiner/parrot/registry"
	"strings"
)

const DingTypeNameInVar = "ding"

const DingNameInVar = "ding"

const AlarmText = `{
    "msgtype": "text", 
    "text": {
        "content": "%s"
    }
}`

type IComponentDing interface {
	GetDingReport() (d DingMessage)
}

type DingMessage interface {
	Text(v, sessionID string, atMobiles []string, isAtAll bool) error
	//Link() error
	//MarkDown() error
	////ActionCard() error
	//FeedCard() error
}

type StandardDing struct {
	c      IContainer
	client *http.HTTPClient
}

func NewStandardDing(c IContainer) *StandardDing {
	client, err := http.NewHTTPClient()
	if err != nil {
		panic(err)
	}
	return &StandardDing{c: c, client: client}
}

func (s *StandardDing) GetDingReport() (d DingMessage) {
	return s
}

func (s *StandardDing) Text(v, sessionID string, atMobiles []string, isAtAll bool) error {
	cacheConf, err := s.c.GetVarConf(DingTypeNameInVar, DingNameInVar)
	if err != nil {
		return fmt.Errorf("%s %v", registry.Join("/", s.c.GetPlatName(), "var", DingTypeNameInVar, DingNameInVar), err)
	}

	context, status, err := s.client.Request(
		"POST",
		cacheConf.GetString("webhook"),
		fmt.Sprintf(AlarmText, "平台:"+s.c.GetPlatName()+"系统:"+s.c.GetServerName()+"\n"+v+
			fmt.Sprintf("\n点击查看详情: %s/%s/%s/%s", cacheConf.GetString("monitor"), s.c.GetPlatName(), strings.Split(s.c.GetServerName(), ".")[0], sessionID)),
		"utf-8",
		map[string]string{
			"Content-Type": "application/json",
		},
	)
	if err != nil {
		return err
	}
	if status != 200 {
		return fmt.Errorf("status:%d context:%s", status, context)
	}
	fmt.Println("context: ", context)
	return nil
}
