package model

import myerr "go_blog/internal/pkg/errors"

type Blog struct {
	BaseModel
	Title   string `gorm:"column:title;type:varchar(128);not null;default:'无标题'" json:"title"`
	Content string `gorm:"column:content;type:mediumtext;not null;" json:"content"`
	Author  uint   `gorm:"column:author;type:int;not null;default:0" json:"author"`
	Visits  uint16 `gorm:"column:visits;type:int unsigned;not null;default:0" json:"visits"`
}

func NewBlog() (*Blog, error) {
	blog := Blog{}
	err := blog.CreateTable()
	if err != nil {
		return &blog, err
	}
	return &blog, err
}
func (blog *Blog) GetBlogById(id uint) *Blog {
	if err := blog.DB().First(blog, id).Error; err != nil {
		return nil
	}
	return blog
}
func (u *Blog) CreateTable() error {
	err := u.DB().AutoMigrate(u)
	if err != nil {
		return err
	}
	return nil
}

func (b *Blog) InsertOne(title string, content []byte, author uint, visits uint16) error {
	b.Title = title
	b.Author = author
	b.Visits = visits
	b.Content = string(content)
	author_id, _ := NewAdminUsers()
	auth := author_id.GetUserById(author)
	if auth != nil {
		b.ID = auth.ID
	} else {
		return myerr.NewBusinessError(myerr.UserDoesNotExist)
	}
	result := b.DB().Create(b)
	return result.Error
}
