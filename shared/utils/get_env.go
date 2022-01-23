package utils

import (
	"github.com/kelseyhightower/envconfig"
)

type env struct {
	DbName           string `split_words:"true"`
	DbPort           string `split_words:"true"`
	DbMasterHost     string `split_words:"true"`
	DbMasterUser     string `split_words:"true"`
	DbMasterPassword string `split_words:"true"`
	DbMasterLogMode  string `split_words:"true"`
	DbReadHost       string `split_words:"true"`
	DbReadUser       string `split_words:"true"`
	DbReadPassword   string `split_words:"true"`
	DbReadLogMode    string `split_words:"true"`
	RootDir          string `split_words:"true"`
	GoPath           string `envconfig:"gopath"`
	AppLoggerLevel   string `split_words:"true"`
	AppLoggerOutput  string `split_words:"true"`
	AppLoggerFormat  string `split_words:"true"`
}

// Env variable
var Env env

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		panic(ErrorsWrap(err, " cannot parse envs"))
	}
	goPath = Env.GoPath
}
