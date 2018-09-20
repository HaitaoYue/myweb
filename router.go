package main

import (
	"github.com/HaitaoYue/myweb/book"
	"github.com/HaitaoYue/myweb/users"
	"github.com/HaitaoYue/myweb/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"net/http"
)

func SetupRouter(r *gin.Engine) {

	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.GET("/foo", func(c *gin.Context) {
		log.Printf("%s in foo %s", utils.Green, utils.Reset)
		c.JSON(http.StatusOK, gin.H{"data": "success", "success": true})
	})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", book.BookableDate)
	}
	r.GET("/bookable", book.GetBookable)

	r.GET("/user/:name", users.UserRetrieve)

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		log.Print(gin.AuthUserKey)
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if err := c.Bind(&json); err == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
}
