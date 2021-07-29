package configs

import "github.com/tkanos/gonfig"

type ConfMysql struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func GetConfMysql() ConfMysql {
	conf := ConfMysql{}
	gonfig.GetConf("configs/confmysql.json", &conf)
	return conf
}
