package models

type User struct {
	Username string `gorm:"uniqueIndex" form:"username" validate:"required,alphanum,min=3,max=20"`
	Password string `form:"password" validate:"required,min=6"`
	Admin    bool
}
