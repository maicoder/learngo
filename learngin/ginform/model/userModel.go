package model

type UserModel struct {
	Email         string `form:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again"`
}
