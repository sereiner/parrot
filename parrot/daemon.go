// Example of a daemon with echo service
package parrot

import (
	"os"

	"github.com/sereiner/parrot/conf/creator"
	"github.com/sereiner/parrot/registry"
	"github.com/urfave/cli"
)

func (m *MicroApp) startAction(c *cli.Context) (err error) {
	msg, err := m.service.Start()
	if err != nil {
		m.xlogger.Error(err)
		return err
	}
	m.xlogger.Info(msg)
	return nil
}
func (m *MicroApp) stopAction(c *cli.Context) (err error) {
	msg, err := m.service.Stop()
	if err != nil {
		m.xlogger.Error(err)
		return err
	}
	m.xlogger.Info(msg)
	return nil
}

func (m *MicroApp) registryAction(c *cli.Context) (err error) {
	if err = m.checkInput(c); err != nil {
		cli.ErrWriter.Write([]byte("  " + err.Error() + "\n\n"))
		cli.ShowCommandHelp(c, c.Command.Name)
		return err
	}
	if err = m.install(); err != nil {
		m.xlogger.Error(err)
		return err
	}
	return nil
}
func (m *MicroApp) serviceAction(c *cli.Context) (err error) {
	if err = m.checkInput(c); err != nil {
		cli.ErrWriter.Write([]byte("  " + err.Error() + "\n\n"))
		cli.ShowCommandHelp(c, c.Command.Name)
		return err
	}
	//安装配置文件
	msg, err := m.service.Install(os.Args[2:]...)
	if err != nil {
		m.xlogger.Error(err)
		return err
	}
	m.xlogger.Info(msg)
	return nil
}
func (m *MicroApp) installAction(c *cli.Context) (err error) {
	if err = m.checkInput(c); err != nil {
		cli.ErrWriter.Write([]byte("  " + err.Error() + "\n\n"))
		cli.ShowCommandHelp(c, c.Command.Name)
		return err
	}

	if err = m.install(); err != nil {
		m.xlogger.Error(err)
		return err
	}

	//安装配置文件
	msg, err := m.service.Install(os.Args[2:]...)
	if err != nil {
		m.xlogger.Error(err)
		return err
	}
	m.xlogger.Info(msg)
	return nil
}
func (m *MicroApp) removeAction(c *cli.Context) (err error) {
	msg, err := m.service.Remove()
	if err != nil {
		m.xlogger.Error(err)
		return err
	}
	m.xlogger.Info(msg)
	return nil
}
func (m *MicroApp) statusAction(c *cli.Context) (err error) {
	msg, err := m.service.Status()
	if err != nil {
		m.xlogger.Error(err)
		return err
	}
	m.xlogger.Info(msg)
	return nil
}

func (m *MicroApp) versionAction(c *cli.Context) (err error) {
	m.xlogger.Debug("git version: ", VERSION)
	m.xlogger.Debug("build time: ", BuildTime)
	return nil
}

func (m *MicroApp) install() (err error) {
	m.logger.PauseLogging()
	defer m.logger.StartLogging()
	//创建注册中心
	rgst, err := registry.NewRegistryWithAddress(m.RegistryAddr, m.logger)
	if err != nil {
		return err
	}

	//自动创建配置
	creator := creator.NewCreator(m.PlatName, m.SystemName, m.ServerTypes, m.ClusterName, m.Conf,
		m.RegistryAddr, rgst,
		m.xlogger)
	err = creator.Start()
	if err != nil {
		return err
	}
	return nil

}
