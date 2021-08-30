package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static") // 静态资源
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"sddddddddddd": "22"})

	})
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"mag":"Get",
		})
	})
	r.POST("/book/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"msg":"POST"})
		
	})
	r.PUT("/book/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"id": "222"})
	})
	r.DELETE("book/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"id": "DELETE"})
	})


	r.GET("posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK,"posts/index.html",gin.H{
			"title":"posts/index",
		})
	})
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("index.tmpl")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})
	r.Run(":7070")

}
