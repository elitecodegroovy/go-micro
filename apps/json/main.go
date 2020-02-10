package main

import (
	"fmt"
	l "github.com/elitecodegroovy/goutil/logger"
	cfg "github.com/micro/go-micro/v2/apps/json/basic/config"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/json"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	"go.uber.org/zap"
	"time"
)

var (
	jsonConf = "conf.json"
	ymlConf  = "app.yml"
	//application name
	appName        = "app"
	userModuleName = "auth_srv"

	//logger object
	log = l.GetLogger()
	//application configuration parameters
	appCfg = &cfg.AppCfg{}
)

type Host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func loadJsonFile() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal("....error in func config.NewConfig()" + err.Error())
	}

	enc := json.NewEncoder()

	// Load toml file with encoder
	conf.Load(file.NewSource(
		file.WithPath("./apps/json/conf/"+file.DefaultPath),
		source.WithEncoder(enc),
	))
	var host Host

	go func() {

		log.Info("[init] 侦听配置变动 ...")

		// 开始侦听变动事件
		watcher, err := conf.Watch()
		if err != nil {
			log.Fatal(err.Error())
		}

		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatal(err.Error())
			}

			log.Info("[init] 侦听配置变动: " + string(v.Bytes()))
			conf.Get("hosts", "database").Scan(&host)
			fmt.Println(host.Address, host.Port)

		}
	}()

	conf.Get("hosts", "database").Scan(&host)
	fmt.Println(host.Address, host.Port)
}

func loadYmlFile() {
	cfg.LoadConfigurationFile([]string{ymlConf})

	cfg.SetAppName(appName)
	cfg.GetConfigurator().Path(userModuleName, appCfg)

	log.Info(" loaded the app configuration ,0",
		zap.String("app conf",
			fmt.Sprintf("app: %s, version:%s, address: %s, port: %d",
				appCfg.Name, appCfg.Version, appCfg.Address, appCfg.Port)))
}

func main() {
	loadJsonFile()
	loadYmlFile()

	time.Sleep(10 * time.Minute)
}
