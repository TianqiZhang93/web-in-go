/*
modification history
--------------------
2020/05/14 15:45:20, by NKztq, create
*/

// Package main is special.  It defines a
// standalone executable program, not a library.
// Within package main the function main is also
// special-it's where execution of the program begins.
// Whatever main does is what the program does.
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "go api test",
		})
	})
	r.Run(":8001")
}
