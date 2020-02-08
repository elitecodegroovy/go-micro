package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/json"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/util/log"
)

var (
	jsonConf = "conf.json"
)

type Host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func loadingJsonFile() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal("....error in func config.NewConfig()", err.Error())
	}

	enc := json.NewEncoder()

	// Load toml file with encoder
	conf.Load(file.NewSource(
		file.WithPath("./apps/json/conf/"+jsonConf),
		source.WithEncoder(enc),
	))

	var host Host
	conf.Get("hosts", "database").Scan(&host)

	fmt.Println(host.Address, host.Port)
}

func main() {
	loadingJsonFile()
}
