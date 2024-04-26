package main

import (
	"net/http"

	"it-post/database"
	"it-post/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectToDB()

    r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	r.Use(cors.New(config))
    
    // 查詢全部的 post
	var posts []model.Post
	database.DB.Find(&posts)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"posts": posts,
		})
	})

	r.POST("/posts", func(c *gin.Context) {
		var posts model.Post
		if err := c.ShouldBindJSON(&posts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad request data",
			})
			return
		}

		if err := database.DB.Create(&posts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create post",
			})
			return

		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Create post successfully",
		})
	})

	r.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		var post model.Post
		
		if err := database.DB.First(&post, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Failed to fetch posts",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"post":	post,
		})
	})

	r.PUT("/posts/:id", func (c *gin.Context) {
		id := c.Param("id")
		var post model.Post

		if err := database.DB.First(&post, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Failed to fetch posts",
			})	
			return
		}

		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad request data",
			})
			return
		} 

		if err := database.DB.Save(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update post",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Post updated successfully",
		})
	})

	r. DELETE("posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := database.DB.Delete(&model.Post{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Post not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{ "message": "Delete post successfully",})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}