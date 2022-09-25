package form

type BlogForm struct {
	Author  uint   `form:"author" json:"author" binding:"required,min=1"`
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Visits  uint16 `form:"visits" json:"visits" binding:"required"`
}

func NewBlogForm() *BlogForm {
	return &BlogForm{}
}
