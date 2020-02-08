package config

import (
	"github.com/micro/go-micro/v2/config/source"
	"go.uber.org/zap"
	"strconv"
)

// AppCfg common config
type AppCfg struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func (a *AppCfg) Addr() string {
	return a.Address + ":" + strconv.Itoa(a.Port)
}

type Options struct {
	Apps    map[string]interface{}
	AppName string
	Sources []source.Source
}

type Etcd struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

type ZapCfg struct {
	zap.Config
	LogFileDir    string `json:logFileDir`
	AppName       string `json:"appName"`
	ErrorFileName string `json:"errorFileName"`
	WarnFileName  string `json:"warnFileName"`
	InfoFileName  string `json:"infoFileName"`
	DebugFileName string `json:"debugFileName"`
	MaxSize       int    `json:"maxSize"` // megabytes
	MaxBackups    int    `json:"maxBackups"`
	MaxAge        int    `json:"maxAge"` // days
}

type Option func(o *Options)

func WithSource(src source.Source) Option {
	return func(o *Options) {
		o.Sources = append(o.Sources, src)
	}
}

func WithApp(appName string) Option {
	return func(o *Options) {
		o.AppName = appName
	}
}
