package main

import (
	"github.com/real013228/user-segment-service/initializers"
	"github.com/real013228/user-segment-service/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Segment{})
}
