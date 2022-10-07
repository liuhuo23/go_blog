package model

import (
	uuid "github.com/satori/go.uuid"
	"go_blog/internal/pkg/errors"
	log "go_blog/internal/pkg/logger"
	"go_blog/internal/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	ID       uint64           `gorm:"column:id;not null;primarykey;AUTO_INCREMENT" json:"id"`
	CreateAt utils.FormatDate `gorm:"column:created_id;type:timestamp;<-:create" json:"create_at"`
	UpdateAt utils.FormatDate `gorm:"column:update_at;type:timestamp" json:"update_at"`
	Title    string           `gorm:"column:title;type:varchar(128);not null;default:'无标题'" json:"title"`
	Content  string           `gorm:"column:content;type:mediumtext;not null;" json:"content"`
	Author   uuid.UUID        `gorm:"column:author;type:varchar(50)" json:"author"`
	Visits   uint16           `gorm:"column:visits;type:int unsigned;not null;default:0" json:"visits"`
	Comment  uint64           `gorm:"column:comment;not null;default:0"  json:"comment"`
	Read     uint64           `gorm:"column:read;not null;default:0" json:"read"`
	Like     uint64           `gorm:"column:like;not null;default:0" json:"like"`
}

func NewBlog() (*Blog, error) {
	blog := Blog{}
	err := blog.CreateTable()
	if err != nil {
		return &blog, err
	}
	return &blog, err
}

// DB 获取DB实例
func (blog *Blog) DB() *gorm.DB {
	return DB()
}

// BeforeCreate 创建前的回调
func (blog *Blog) BeforeCreate(tx *gorm.DB) error {
	blog.CreateAt = utils.FormatDate{Time: time.Now()}
	return nil
}

// AfterCreate 插入后的回调，如果不满足条件 则回滚插入
func (blog *Blog) AfterCreate(tx *gorm.DB) error {
	admin, _ := NewAdminUsers()
	log.Logger.Sugar().Info("%+v", blog)
	err := admin.DB().First(&admin, blog.Author).Error
	if err != nil {
		log.Logger.Sugar().Info("crate blog err:", err.Error())
		return errors.NewBusinessError(errors.UserDoesNotExist, "该用户不存在！")
	}
	return nil
}

// GetBlogById 通过博客id获取博客
func (blog *Blog) GetBlogById(id uint) *Blog {
	if err := blog.DB().First(blog, id).Error; err != nil {
		return nil
	}
	return blog
}

// CreateTable 创建表
func (u *Blog) CreateTable() error {
	err := u.DB().AutoMigrate(u)
	if err != nil {
		return err
	}
	return nil
}

// InsertOne 添加一个博客
func (b *Blog) InsertOne(title string, content []byte, author string, visits uint16) error {
	b.Title = title
	b.Visits = visits
	b.Content = string(content)
	b.Author, _ = uuid.FromString(author)
	result := b.DB().Create(b)
	return result.Error
}

// GetList 获取所有数据
func (b *Blog) GetList(page int, limit int) []Blog {
	var blogs []Blog
	b.DB().Offset((page - 1) * limit).Limit(limit).Find(&blogs)
	return blogs
}
