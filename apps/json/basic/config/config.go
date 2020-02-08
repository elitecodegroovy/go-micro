package config

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/yaml"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/util/log"
	"go.uber.org/zap"
	"os"
	"sync"
)

var (
	m      sync.RWMutex
	inited bool
	// 默认配置器
	c = &configurator{}
)

// Configurator 配置器
type Configurator interface {
	App(name string, config interface{}) (err error)
	Path(path string, config interface{}) (err error)
}

// configurator 配置器
type configurator struct {
	conf    config.Config
	appName string
}

// Init 初始化配置
func init() {
	log.Logf("> config init ...")
	opts := Options{}
	c.init(opts)
}

func (c *configurator) App(name string, config interface{}) (err error) {

	v := c.conf.Get(name)
	if v != nil {
		log.Info(">>>", zap.String("auth_srv", fmt.Sprintf("%s", string(v.Bytes()))))
		err = v.Scan(config)
	} else {
		err = fmt.Errorf("[App] 配置不存在，err：%s", name)
	}

	return
}

func (c *configurator) Path(path string, config interface{}) (err error) {
	v := c.conf.Get(c.appName, path)
	if v != nil {
		err = v.Scan(config)
	} else {
		err = fmt.Errorf("[Path] 配置不存在，err：%s", path)
	}

	return
}

// c 配置器
func GetConfigurator() Configurator {
	return c
}

func GetC() *configurator {
	return c
}

func (c *configurator) init(ops Options) (err error) {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[init] 配置已经初始化过")
		return
	}
	c.conf, err = config.NewConfig()
	c.appName = ops.AppName

	// 加载配置
	err = c.conf.Load(ops.Sources...)
	if err != nil {
		log.Fatal(err)
	}
	// 标记已经初始化
	inited = true
	return
}

func SetAppName(appName string) {
	c.appName = appName
}

func LoadConfigurationFile(configurationFileNames []string) {
	encode := yaml.NewEncoder()
	for _, app := range configurationFileNames {
		if err := c.conf.Load(file.NewSource(
			file.WithPath("/home/app/goapp/src/github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/auth/conf/"+app+".yml"),
			source.WithEncoder(encode),
		)); err != nil {
			log.Fatal("[loadAndWatchConfigFile] 加载应用配置文件 异常，%s", zap.String("err:", err.Error()))
			os.Exit(1)
		}
	}
}

func GetCfgValueByName(appName string) []byte {
	return c.conf.Get(appName).Bytes()
}
