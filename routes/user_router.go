package routes

import (
	handler "BlogAPI/handler"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/users")
	{
		userRouter.POST("/", handler.CreateUser)
		userRouter.POST("/login", handler.Login)
		userRouter.POST("/getuser", handler.GetUserByUsername)
	}
}
