package controller

import (
	"User/models"
	"User/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo *repositories.UserRepository
}

func NewUserController(repo *repositories.UserRepository) *UserController {
	return &UserController{Repo: repo}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.Repo.CreateNewUser(&user)
	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully Created!"})
}

func (ctrl *UserController) GetAllUser(c *gin.Context) {
	result, row_count, error := ctrl.Repo.GetAllUsers()

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There is a problem database"})
		return
	}

	if row_count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not founded any users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All Users", "data": result})
}
