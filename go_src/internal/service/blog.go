package service

import (
	uuid "github.com/satori/go.uuid"
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

func (b *BlogServer) CreateBlog(title string, content []byte, author string, visits uint16) (uuid.UUID, error) {
	blog, berr := model.NewBlog()
	if berr != nil {
		log.Logger.Sugar().Error("创建blog失败，err：", berr.Error())
		return uuid.Nil, berr
	}
	err := blog.InsertOne(title, content, author, visits)
	if err != nil {
		log.Logger.Sugar().Error("插入数据失败！err:", err)
		return uuid.Nil, err
	}
	var id []uuid.UUID
	blog.DB().Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)
	log.Logger.Sugar().Info(id)
	return id[0], nil
}

func (b *BlogServer) GetList(page int, limit int) []model.Blog {
	blog, _ := model.NewBlog()
	data := blog.GetList(page, limit)
	return data
}
