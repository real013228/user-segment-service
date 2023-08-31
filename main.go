package main

import (
	"github.com/gin-gonic/gin"
	"github.com/real013228/user-segment-service/controllers"
	"github.com/real013228/user-segment-service/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func setUserControllers(r *gin.Engine) {
	r.POST("/user/create", controllers.CreateUser)
	r.GET("/user/get/all", controllers.GetUsers)
	r.GET("/user/get/:id", controllers.GetUserByID)
	r.PUT("/user/update/:id", controllers.UpdateUser)
	r.DELETE("/user/delete/:id", controllers.DeleteUser)
	r.PUT("/user/register", controllers.RegisterUser)
}

func setSegmentControllers(r *gin.Engine) {
	r.POST("/segment/create", controllers.CreateSegment)
	r.GET("/segment/get/all", controllers.GetSegments)
	r.GET("/segment/get/:id", controllers.GetSegmentById)
	r.DELETE("/segment/delete/:id", controllers.DeleteSegment)
}

func main() {
	r := gin.Default()
	setUserControllers(r)
	setSegmentControllers(r)
	r.Run()
}
