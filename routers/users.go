package routers

import (
	"fazz/backend/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.ListUsers)
	routerGroup.GET("/:id", controllers.DetailUser)
	routerGroup.POST("/", controllers.CreateUser)
	routerGroup.PATCH("/:id", controllers.UpdateUser)
	routerGroup.DELETE("/:id", controllers.DeleteUser)
}
