package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})


	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})





	router.POST("/post", func(c *gin.Context) {

		id := c.Query("postid")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)

		c.JSON(200, gin.H{
			"status":  "posted",
			"id": id,
			"page":page,
			"name":name,
			"message":message,

		})
	})

	type postForm1 struct {
		Id     string `form:"id" binding:"required"`
		Page string `form:"page" binding:"required"`
	}

	router.POST("/post1", func(c *gin.Context) {
		var form postForm1
		if c.Bind(&form) == nil{
			c.JSON(200,gin.H{
				"status":200,
				"id":form.Id,
				"page":form.Page,
			})
		}
		fmt.Printf("id: %s ", form.Id)
	})

	type LoginForm struct {
		User     string `form:"user" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	router.POST("/login", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:
		// c.BindWith(&form, binding.Form)
		// or you can simply use autobinding with Bind method:
		var form LoginForm
		// in this case proper binding will be automatically selected
		if c.Bind(&form) == nil {
			if form.User == "1001" && form.Password == "1002" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	router.Run(":8080") // listen and server on 0.0.0.0:8080



}
