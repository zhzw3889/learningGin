package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Get
	r.GET("/get", getMsg)

	// Post
	r.POST("/post", postMsg)

	r.GET("/testRedirect", testRedirect)
	// redirect to outside
	r.GET("/redirect1", redirectToOut)
	// redirect another path
	r.GET("/redirect2", func(c *gin.Context) {
		c.Request.URL.Path = "/testRedirect"
		r.HandleContext(c)
	})

	r.GET("/getOtherData", func(ctx *gin.Context) {
		url := "https://img2.baidu.com/it/u=1649006510,1140670358&fm=253&fmt=auto&app=138&f=JPEG?w=600&h=350"
		response, err := http.Get(url)
		if err != nil || response.StatusCode != http.StatusOK {
			ctx.Status(http.StatusServiceUnavailable)
			return
		}
		body := response.Body
		contentlength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		// 转发的方法
		ctx.DataFromReader(http.StatusOK, contentlength, contentType, body, nil)
	})

	r.Run(":9090")

}

// get
func getMsg(c *gin.Context) {
	name := c.Query("name")
	// c.String(http.StatusOK, "欢迎你，%s", name)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "get reponse",
		"data": "Welcome, " + name,
	})
}

// post
func postMsg(c *gin.Context) {
	// name := c.DefaultPostForm("name", "default: gin")
	json := make(map[string]interface{})
	c.BindJSON(&json)
	c.JSON(http.StatusOK, gin.H{
		"name":     json["name"],
		"password": json["password"],
		"msg":      "post reponse",
	})
}

func testRedirect(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "redirect reponse",
	})
}

// 重定向到外部网络
func redirectToOut(c *gin.Context) {
	url := "http://www.baidu.com"
	c.Redirect(http.StatusMovedPermanently, url)
}
