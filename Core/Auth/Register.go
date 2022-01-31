package auth

type Register struct {
	UserName string `form:"username" binding:"alpha,required"`
	Email    string `form:"email" binding:"email,required"`
	Password string `form:"password" binding:"required"`
}
