package config

import (
	"gopkg.in/ini.v1"
)

type config struct {
	App   app   `ini:"app"`
	Redis redis `ini:"redis"`
	DB    db    `ini:"database"`
	Log   log   `ini:"log"`
	Mini  mini  `ini:"mini"`
}

var (
	App   app
	Redis redis
	DB    db
	Log   log
	Mini  mini
)

func LoadConfig() {
	cfg, err := ini.Load(".env", ".env.local")
	if err != nil {
		panic("load config err:" + err.Error())
	}
	c := new(config)
	err = cfg.MapTo(c)
	if err != nil {
		panic("map config err:" + err.Error())
	}
	App = c.App
	Redis = c.Redis
	DB = c.DB
	Log = c.Log
	Mini = c.Mini
}
