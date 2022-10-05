package model

import uuid "github.com/satori/go.uuid"

type Blog struct {
	BaseModel
	Title      string     `gorm:"column:title;type:varchar(128);not null;default:'无标题'" json:"title"`
	Content    string     `gorm:"column:content;type:mediumtext;not null;" json:"content"`
	Author     uuid.UUID  `gorm:"column:author;type:varchar(50)" json:"author"`
	Visits     uint16     `gorm:"column:visits;type:int unsigned;not null;default:0" json:"visits"`
	AdminUsers AdminUsers `gorm:"ForeignKey:ID;AssociationForeignKey:Author"`
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

func (b *Blog) InsertOne(title string, content []byte, author string, visits uint16) error {
	b.Title = title
	b.Visits = visits
	b.Content = string(content)
	b.Author, _ = uuid.FromString(author)
	result := b.DB().Create(b)
	return result.Error
}
