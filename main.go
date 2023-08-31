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
	r.POST("/user/create", controllers.CreateUser)                          // create user
	r.GET("/user/get/all", controllers.GetUsers)                            // get all users
	r.GET("/user/get/:id", controllers.GetUserByID)                         // get user by id
	r.PUT("/user/update/:id", controllers.UpdateUser)                       // update user's name
	r.DELETE("/user/delete/:id", controllers.DeleteUser)                    // delete user by id
	r.PUT("/user/register", controllers.RegisterUser)                       // register user to the segment, takes id of user and segment
	r.PUT("/user/register/bunch", controllers.RegisterUserMultipleSegments) // register user to the segments, takes id of user and segment
}

func setSegmentControllers(r *gin.Engine) {
	r.POST("/segment/create", controllers.CreateSegment)       // create segment
	r.GET("/segment/get/all", controllers.GetSegments)         // get all segments
	r.GET("/segment/get/:id", controllers.GetSegmentById)      // get segment by id
	r.DELETE("/segment/delete/:id", controllers.DeleteSegment) // delete segment by id
}

func main() {
	r := gin.Default()
	setUserControllers(r)
	setSegmentControllers(r)
	r.Run()
}
