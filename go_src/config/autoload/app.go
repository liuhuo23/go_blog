package autoload

import (
	"go_blog/pkg/convert"
	"go_blog/pkg/utils"
)

type AppConfig struct {
	AppEnv         string `ini:"app" yaml:"app"`
	Debug          bool   `ini:"debug" yaml:"debug"`
	Language       string `ini:"language" yaml:"language"`
	StaticBasePath string `ini:"base_path" yaml:"base_path"`
}

var App = AppConfig{
	AppEnv:         "local",
	Debug:          true,
	Language:       "zh_CN",
	StaticBasePath: getDefaultPath(),
}

func getDefaultPath() (path string) {
	path = utils.GetRunPath()
	path = convert.GetString(utils.If(path != "", path, "/tmp"))
	return
}
