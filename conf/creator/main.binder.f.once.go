package creator

import "github.com/sereiner/parrot/conf"

type IOnceBinder interface {
	SetMain(*conf.CronServerConf)
	SetTasks(*conf.Tasks)
	imainBinder
}

type OnceBinder struct {
	*mainBinder
}

func newOnceBinder(params map[string]string, inputs map[string]*Input) *OnceBinder {
	return &OnceBinder{
		mainBinder: newMainBinder(params, inputs),
	}
}
func (b *OnceBinder) SetMain(c *conf.CronServerConf) {
	b.mainBinder.SetMainConf(c)
}
func (b *OnceBinder) SetTasks(c *conf.Tasks) {
	b.mainBinder.SetSubConf("task", c)
}
