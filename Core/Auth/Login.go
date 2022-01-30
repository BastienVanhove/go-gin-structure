package auth

type Login struct {
	UserName string `form:"username" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
