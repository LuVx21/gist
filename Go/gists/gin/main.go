package main

import (
    "github.com/gin-gonic/gin"
    //"net/http"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "Blog":   "www.flysnow.org",
            "wechat": "flysnow_org",
        })
    })
    r.Run(":8090")
}

//func main() {
//	r := gin.Default()
//	r.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK, "Who are you?")
//	})
//	r.GET("/user/:name", func(c *gin.Context) {
//		name := c.Param("name")
//		c.String(http.StatusOK, "Hello %s", name)
//	})
//	// 匹配users?name=xxx&role=xxx
//	r.GET("/users", func(c *gin.Context) {
//		name := c.Query("name")
//		role := c.DefaultQuery("role", "teacher")
//		c.String(http.StatusOK, "%s is a %s", name, role)
//	})
//	r.Run(":8090")
//}

//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong1",
//		})
//	})
//	r.Run(":8090")
//}
