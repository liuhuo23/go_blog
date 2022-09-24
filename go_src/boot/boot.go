package boot

import (
	"flag"
	"fmt"
	"go_blog/config"
	"go_blog/data"
	"go_blog/internal/pkg/logger"
	"go_blog/internal/routers"
	"go_blog/internal/validator"
)

const version = "0.0.1"

func init() {
	var (
		configPath   string
		printVersion bool
	)
	flag.StringVar(&configPath, "c", "", "请输入配置文件的绝对路径")
	flag.BoolVar(&printVersion, "v", false, "查看版本")
	flag.Parse()

	if printVersion {
		println(version)
	}

	//1. 初始化配置
	config.InitConfig(configPath)
	// 2.初始化zap日志
	logger.InitLogger()
	// 3. 初始化数据库
	data.InitData()
	// 4. 初始化验证器
	validator.InitValidatorTrans("zh")
}
func Run() {
	r := routers.SetRouters()
	err := r.Run(fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port))
	if err != nil {
		panic(err)
	}
}
