package router

import (
	"example/web-service-gin/pkg/user"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const dataSourceName = "root:@tcp(localhost:3306)/user_management"

// SetupRouter configures the API routes
func SetupRouter(userService *user.UserService) *gin.Engine {
	router := gin.Default()

	// user-related API endpoints
	userRoutes := router.Group("/users")
	{
		userController := user.NewUserController(userService)

		userRoutes.POST("", userController.CreateUserHandler)
		userRoutes.GET("", userController.GetAllUsersHandler)
		userRoutes.GET("/:id", userController.GetUserHandler)
		userRoutes.DELETE("/:id", userController.DeleteUserHandler)
	}

	return router
}
