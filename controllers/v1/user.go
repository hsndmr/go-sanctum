package controllerv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/controllers/middleware"
	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/repositories"
)

// RegisterUserRoutes registers the user routes
func RegisterUserRoutes(r *gin.RouterGroup) {
	user := &User{}
	r.POST("/user", user.Create)
	r.POST("/auth/login", user.Login)
	r.GET("/user", middleware.Sanctum(), user.GetUser)
}

// validations
type CreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type User struct {}

// Create creates a new user
func (u *User) Create(c *gin.Context) {
		var createRequest CreateRequest
		if err := c.ShouldBindJSON(&createRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		userRepository := app.C.Repository.User

		_, isExistedUser := userRepository.FindByEmail(createRequest.Email)

		if isExistedUser == nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "email already exists"})
			return
		}
		
		user, err := userRepository.Create(&repositories.CreateUserDto{
			Name: createRequest.Name,
			Email: createRequest.Email,
			Password: createRequest.Password,
		})

		if err !=nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		personalAccessTokenRepository := app.C.Repository.PersonalAccessToken

		_, token, err := personalAccessTokenRepository.Create(&repositories.CreatePersonalAccessTokenDto{
			Name: "Personal Access Token",
			User: user,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": gin.H{
				"user": user,
				"token": token,
			},
		})
}

func (u *User) Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := app.C.Repository.User.FindByEmail(loginRequest.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email or password is not correct"})
		return
	}

	if !app.C.Service.Crypto.CheckPasswordHash(loginRequest.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email or password is not correct"})
		return
	}

	_, token, err := app.C.Repository.PersonalAccessToken.Create(&repositories.CreatePersonalAccessTokenDto{
		User: user,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": user,
			"token": token,
		},
	})
}

func (u *User) GetUser(c *gin.Context) {
	user := c.MustGet("user").(*ent.User)
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": user,
		},
	})
}