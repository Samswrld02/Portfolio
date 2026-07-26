package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Samswrld02/Portfolio/dto"
	"github.com/Samswrld02/Portfolio/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	Db *gorm.DB
}

// create claims
type JwtAuthClaims struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}

func (auth *Auth) Login(c *echo.Context) error {
	var user models.User
	var userDTO dto.User

	//bind to user dto
	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf(" err: %s", err.Error())})
	}

	//validate dto
	if err := c.Validate(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("validation err: %s", err.Error())})
	}

	//make request to backend
	if err := auth.Db.Where("username = ?", userDTO.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user does not exist"})
	}

	//model password is hashed, validate user
	if err := auth.validateHash(&user.Password, &userDTO.Password); err != nil {
		return echo.ErrUnauthorized
	}

	//create claims
	claims := &JwtAuthClaims{
		userDTO.Username,
		false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((time.Hour * 1))),
			Issuer:    "sam the man",
			Subject:   "user",
		},
	}

	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	c.Response().Header().Set("Authorization", "Bearer: "+t)

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// registration handler
func (auth *Auth) Register(c *echo.Context) error {
	var UserModel models.User
	var userDTO dto.CreateUserDTO

	//bind request data and validate
	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Err:": err.Error()})
	}

	if err := c.Validate(&userDTO); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Err:": err.Error()})
	}

	//map values
	UserModel.Username = userDTO.Username
	hash, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "something went wrong try again later")
	}
	UserModel.Password = string(hash)

	//create user
	if err := auth.Db.Create(&UserModel).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Err:": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"succes:": "user created"})
}

// validate hash
func (auth *Auth) validateHash(hashedPassword *string, password *string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(*password)); err != nil {
		return err
	}
	return nil
}
