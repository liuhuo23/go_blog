package boot

import (
	"flag"
	"go_blog/config"
)
const version = "0.0.1"
func init() {
	var(
		configPath string
		printVersion bool
	)
	flag.StringVar(&configPath, "c", "", "请输入配置文件的绝对路径")
	flag.BoolVar(&printVersion, "v", false, "查看版本")
	flag.Parse()

	if printVersion {
		println(version)
	}
	
	//3. 


}