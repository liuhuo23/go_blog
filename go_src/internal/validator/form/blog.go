package form

type BlogForm struct {
	Author  string `form:"author" json:"author" binding:"required"`
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Visits  uint16 `form:"visits" json:"visits" binding:"required"`
}

func NewBlogForm() *BlogForm {
	return &BlogForm{}
}
