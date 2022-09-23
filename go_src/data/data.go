package data

import (
	c "go_blog/config"
	"sync"
)

var once sync.Once

func InitData() {
	once.Do(func() {
		if c.Config.Mysql.Enable {
			// 初始化
			InitData()
		}
		if c.Config.Redis.Enable {
			//初始化redis
			initRedis()
		}
	})
}
