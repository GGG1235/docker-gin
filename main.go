package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "person"
		}
		c.JSON(200,gin.H{
			"name" : name,
			"ip" : c.ClientIP(),
		})
	})
	_ = router.Run("0.0.0.0:8080")
}
