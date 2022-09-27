package service

import (
	"go_blog/internal/model"
	log "go_blog/internal/pkg/logger"
)

type BlogServer struct {
}

func NewBlogServer() *BlogServer {
	return &BlogServer{}
}

func (b *BlogServer) QueryById(id uint) (*model.Blog, error) {
	//获取blog
	blog, berr := model.NewBlog()
	if berr != nil {
		log.Logger.Sugar().Info("创建表失败，err:", berr.Error())
		return blog, berr
	}
	data := blog.GetBlogById(id)
	return data, nil
}

func (b *BlogServer) CreateBlog(title string, content []byte, author uint, visits uint16) (uint, error) {
	blog, berr := model.NewBlog()
	if berr != nil {
		log.Logger.Sugar().Error("创建blog失败，err：", berr.Error())
		return 0, berr
	}
	err := blog.InsertOne(title, content, author, visits)
	if err != nil {
		return 0, err
	}
	var id []uint
	blog.DB().Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)
	return id[0], nil
}
