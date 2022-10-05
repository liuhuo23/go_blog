package model

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go_blog/data"
	"go_blog/internal/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type BaseModel struct {
	ID       uuid.UUID        `gorm:"column:id;type:varchar(50);not null;primarykey" json:"id"`
	CreateAt utils.FormatDate `gorm:"column:created_id;type:timestamp;<-:create" json:"create_at"`
	UpdateAt utils.FormatDate `gorm:"column:update_at;type:timestamp" json:"update_at"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.NewV4()
	fmt.Println(uuid.NewV4())
	base.CreateAt = utils.FormatDate{Time: time.Now()}
	return nil
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
