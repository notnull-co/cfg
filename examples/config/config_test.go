package config

import (
	"fmt"
	"regexp"
	"time"

	"github.com/notnull-co/cfg"
)

type Config struct {
	App struct {
		Environment string `cfg:"environment" validate:"required"`
	} `cfg:"app"`
	Server struct {
		Host         string        `cfg:"host" default:"0.0.0.0"`
		Port         int           `cfg:"port" default:"80"`
		ReadTimeout  time.Duration `cfg:"read_timeout" default:"30s"`
		WriteTimeout time.Duration `cfg:"write_timeout" default:"30s"`
	} `cfg:"server"`
	Logger struct {
		Level   string         `cfg:"level" default:"info"`
		Pattern *regexp.Regexp `cfg:"pattern" default:".*"`
	} `cfg:"logger"`
	Certificate struct {
		Version    int       `cfg:"version"`
		DNSNames   []string  `cfg:"dns_names" default:"[kkyr,kkyr.io]"`
		Expiration time.Time `cfg:"expiration" validate:"required"`
	} `cfg:"certificate"`
}

func ExampleLoad() {
	var conf Config
	err := cfg.Load(&conf, cfg.TimeLayout("2006-01-02"))
	if err != nil {
		panic(err)
	}

	fmt.Println(conf.App.Environment)
	fmt.Println(conf.Server.Host)
	fmt.Println(conf.Server.Port)
	fmt.Println(conf.Server.ReadTimeout)
	fmt.Println(conf.Server.WriteTimeout)
	fmt.Println(conf.Logger.Level)
	fmt.Println(conf.Logger.Pattern)
	fmt.Println(conf.Certificate.Version)
	fmt.Println(conf.Certificate.DNSNames)
	fmt.Println(conf.Certificate.Expiration.Format("2006-01-02"))

	// Output:
	// dev
	// 0.0.0.0
	// 443
	// 1m0s
	// 30s
	// debug
	// [a-z]+
	// 1
	// [kkyr kkyr.io]
	// 2020-12-01
}
