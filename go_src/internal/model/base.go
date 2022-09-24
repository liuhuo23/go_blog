package model

import (
	"go_blog/data"
	"go_blog/internal/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	ID       uint             `gorm:"column:id;type:int(11) unsigned AUTO_INCREMENT;not null;primarykey" json:"id"`
	CreateAt utils.FormatDate `gorm:"column:created_id;type:timestamp;<-:create" json:"create_at"`
	UpdateAt utils.FormatDate `gorm:"column:update_at;type:timestamp" json:"update_at"`
}

func (model *BaseModel) DB() *gorm.DB {
	return DB()
}

type ContainsDeleteBaseModel struct {
	BaseModel
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11); unsigned;not null;default:0; index" json:"-"`
}

func DB() *gorm.DB {
	return data.MysqlDB
}
