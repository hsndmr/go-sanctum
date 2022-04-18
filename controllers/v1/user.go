package controllerv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/repositories"
)

// RegisterUserRoutes registers the user routes
func RegisterUserRoutes(r *gin.RouterGroup) {
	user := &User{}
	r.POST("/user", user.Create)
}

// validations
type CreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type User struct {}

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

		c.JSON(http.StatusCreated, gin.H{
			"data": user,
		})
}