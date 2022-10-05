package form

type Login struct {
	UserName string `form:"username" json:"username" binding:"required,min=5"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
	Mobile   string `form:"mobile" json:"mobile"  binding:"required"`
}

func LoginForm() *Login {
	return &Login{}
}
