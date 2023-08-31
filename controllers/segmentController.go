package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/real013228/user-segment-service/initializers"
	"github.com/real013228/user-segment-service/models"
)

func CreateSegment(c *gin.Context) {
	// get data off req body
	var body struct {
		Name string
	}

	c.Bind(&body)

	// check there is no segment with the same name
	var segmentCheck models.Segment
	if err := initializers.DB.Where("name = ?", body.Name).First(&segmentCheck).Error; err == nil {
		c.JSON(400, gin.H{
			"kind":    "segment_already_exists",
			"message": "Segment with the same name already exists " + body.Name,
		})
		return
	}

	// create segment
	segment := models.Segment{Name: body.Name}

	res := initializers.DB.Create(&segment)

	if res.Error != nil {
		c.Status(400)
		return
	}

	// respond with it
	c.JSON(200, gin.H{
		"segment": segment,
	})
}

func GetSegments(c *gin.Context) {
	// get the segments
	var segments []models.Segment
	initializers.DB.Find(&segments)

	// respond with them
	c.JSON(200, gin.H{
		"segments": segments,
	})
}

func GetSegmentById(c *gin.Context) {
	// get id off url
	id := c.Param("id")

	// get segment
	var segment models.Segment
	if err := initializers.DB.First(&segment, id).Error; err != nil {
		c.JSON(400, gin.H{
			"kind":    "segment_not_found",
			"message": "There is no segment with id" + id,
		})
		return
	}

	// respond with it
	c.JSON(200, gin.H{
		"segment": segment,
	})
}

func DeleteSegment(c *gin.Context) {
	// get the id off segment
	id := c.Param("id")

	// get the segment with this id
	var segment models.Segment
	if err := initializers.DB.First(&segment, id).Error; err != nil {
		c.JSON(400, gin.H{
			"kind":    "segment_not_found",
			"message": "There is no segment with id" + id,
		})
		return
	}
	// delete the users from segment
	var users []models.User
	initializers.DB.Find(&users)

	for i := 0; i < len(users); i++ {
		initializers.DB.Model(&users[i]).Association("Segments").Delete(segment)
	}

	initializers.DB.Delete(&segment, id)
	c.Status(200)
}
