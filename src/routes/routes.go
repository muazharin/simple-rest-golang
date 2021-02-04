package routes

import (
	c_user "gin-rest-mongo/src/controllers"

	"github.com/gin-gonic/gin"
)

type Routes struct{}

func (c Routes) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/users", c_user.GetAllUser)
		api.POST("/users", c_user.CreateUser)
		api.GET("/users/:id", c_user.GetUser)
		api.PUT("/users/:id", c_user.UpdateUser)
		api.DELETE("/users/:id", c_user.DeleteUser)

		api.GET("/profile", c_user.GetAllProfile)
		api.POST("/profile", c_user.CreateProfile)
		api.DELETE("/profile/:name", c_user.DeleteProfile)
	}

	r.Run()
}
