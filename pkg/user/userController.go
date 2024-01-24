package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *UserService
}

func NewUserController(us *UserService) *UserController {
	return &UserController{UserService: us}
}

func (uc *UserController) GetAllUsersHandler(c *gin.Context) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUserHandler(c *gin.Context) {
	var u User

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uc.UserService.CreateUser(u.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "User created successfully"})
}

func (uc *UserController) GetUserHandler(c *gin.Context) {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	existingUser, err := uc.UserService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingUser)
}

func (uc *UserController) DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	err = uc.UserService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with ID %d deleted", id)})
}
