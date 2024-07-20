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
		//userRouter.GET("/", handler.GetUser)
		/*userRouter.PUT("/", handlers.UpdateUser)
		userRouter.PUT("/password", handlers.UpdatePassword)
		userRouter.PUT("/password/reset", handlers.ResetPassword)
		userRouter.DELETE("/", handlers.DeleteUser)*/
	}
}
