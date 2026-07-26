package dto

type User struct {
	Username string `form:"username" validate:"required,alphanum,min=3,max=20"`
	Password string `form:"password" validate:"required,min=6"`
}

// creation dto
type CreateUserDTO struct {
	Username string `form:"username" validate:"required,alphanum,min=3,max=20"`
	Password string `form:"password" validate:"required,min=6"`
}
