package main

import (
	"fmt"
	"luvx/api/greet"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(greet.Add(1, 3))
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
