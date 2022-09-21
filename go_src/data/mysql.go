package data

import (
	c "go_blog/config"
	"log"

	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

type Writer interface {
	Printf(string, ...interface{})
}

type WriteLog struct{}

func (w WriteLog) Printf(format string, args ...interface{}) {
	if c.Config.Mysql.PrintSql {
		log.Logger.Sugar().Infof(format, args...)
	}
}
