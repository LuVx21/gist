package main

import (
    "fmt"
)

func main() {
    fmt.Println("hello")
    panic("hello")
}

//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong1",
//		})
//	})
//	r.Run(":8090")
//}
