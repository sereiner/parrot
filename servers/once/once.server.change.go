package once

import (
	"fmt"

	"github.com/sereiner/parrot/conf"
)

//SetMetric 重置metric
func (s *OnceServer) SetMetric(metric *conf.Metric) error {
	s.metric.Stop()
	if metric.Disable {
		return nil
	}
	if err := s.metric.Restart(metric.Host, metric.DataBase, metric.UserName, metric.Password, metric.Cron, s.Logger); err != nil {
		err = fmt.Errorf("metric设置有误:%v", err)
		return err
	}
	return nil
}

//StopMetric stop metric
func (s *OnceServer) StopMetric() error {
	s.metric.Stop()
	return nil
}

//SetTasks 设置定时任务
func (s *OnceServer) SetTasks(redisSetting string, tasks []*conf.Task) (err error) {
	s.Processor, err = s.getProcessor(redisSetting, tasks)
	return err
}

//SetTrace 显示跟踪信息
func (s *OnceServer) SetTrace(b bool) {
	s.conf.SetMetadata("show-trace", b)
	return
}
