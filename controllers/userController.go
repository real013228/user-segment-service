package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/real013228/user-segment-service/initializers"
	"github.com/real013228/user-segment-service/models"
)

func CreateUser(c *gin.Context) {
	// get data off req body
	var body struct {
		Name string
	}

	c.Bind(&body)

	// create user
	user := models.User{Name: body.Name}

	res := initializers.DB.Create(&user)

	if res.Error != nil {
		c.Status(400)
		return
	}

	// respond with it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUsers(c *gin.Context) {
	// get the users
	var users []models.User
	initializers.DB.Find(&users)
	for i := range users {
		initializers.DB.Model(&users[i]).Association("Segments").Find(&users[i].Segments)
	}
	// respond with them
	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUserByID(c *gin.Context) {
	// get id off url
	id := c.Param("id")
	// get user
	var user models.User

	// check existing
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(400, gin.H{
			"kind":    "user_not_found",
			"message": "There is no user with id" + id,
		})
		return
	}
	initializers.DB.Model(&user).Association("Segments").Find(&user.Segments)
	// respond with it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	// get the id off url
	id := c.Param("id")

	// get the data
	var body struct {
		Name string
	}

	c.Bind(&body)

	// find the user were updating
	var user models.User

	// check existing
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(400, gin.H{
			"kind":    "user_not_found",
			"message": "There is no user with id" + id,
		})
		return
	}

	// update it
	initializers.DB.Model(&user).Updates(models.User{Name: body.Name})

	// respond with it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	// get the id off url
	id := c.Param("id")

	// check existing
	var user models.User
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(400, gin.H{
			"kind":    "user_not_found",
			"message": "There is no user with id" + id,
		})
		return
	}

	// delete the user
	initializers.DB.Model(&user).Association("Segments").Clear()
	initializers.DB.Delete(&models.User{}, id)

	// respond
	c.Status(200)
}

func RegisterUser(c *gin.Context) {
	// get the body
	var body struct {
		UserId    string
		SegmentId string
	}

	c.Bind(&body)

	// find user and segment
	var user models.User
	var segment models.Segment

	// check existing
	if err := initializers.DB.First(&user, body.UserId).Error; err != nil {
		c.JSON(400, gin.H{
			"kind":    "invalid_parameter_error",
			"message": "There is no user with id " + body.UserId,
		})
		return
	}

	if err := initializers.DB.First(&segment, body.SegmentId).Error; err != nil {
		c.JSON(400, gin.H{
			"kind":    "invalid_parameter_error",
			"message": "There is no segment with id " + body.SegmentId,
		})
		return
	}

	// register user in segment
	initializers.DB.Model(&user).Association("Segments").Append(&segment)
	initializers.DB.Model(&user).Association("Segments").Find(&user.Segments)

	// respond
	c.JSON(200, gin.H{
		"user": user,
	})
}

func RegisterUserMultipleSegments(c *gin.Context) {
	var body struct {
		UserId     string
		SegmentsId []string
	}
	c.Bind(&body)

	// find user
	var user models.User

	// check existing
	if err := initializers.DB.First(&user, body.UserId).Error; err != nil {
		c.JSON(400, gin.H{
			"kind":    "invalid_parameter_error",
			"message": "There is no user with id " + body.UserId,
		})
		return
	}

	for i := 0; i < len(body.SegmentsId); i++ {
		// find segment
		var segment models.Segment

		// check existing
		if err := initializers.DB.First(&segment, body.SegmentsId[i]).Error; err != nil {
			c.JSON(400, gin.H{
				"kind":    "invalid_parameter_error",
				"message": "There is no segment with id " + body.SegmentsId[i],
			})
			return
		}

		// register user in segment
		initializers.DB.Model(&user).Association("Segments").Append(&segment)
		initializers.DB.Model(&user).Association("Segments").Find(&user.Segments)
	}

	// respond
	c.JSON(200, gin.H{
		"user": user,
	})
}
