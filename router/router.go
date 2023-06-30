package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/NKztq/go-web-api/config"
	"github.com/NKztq/go-web-api/middleware"
)

func InitRouter(conf *config.LocalConfig) (*gin.Engine, error) {
	// set gin logger
	accessLogger, err := os.OpenFile(conf.Server.AccessLogPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return nil, fmt.Errorf("os.OpenFile(%s): %s", conf.Server.AccessLogPath, err.Error())
	}

	router := gin.New()
	router.Use(gin.LoggerWithWriter(accessLogger, ""))

	// set whitelist value
	router.Use(func(c *gin.Context) {
		c.Set(middleware.ContextKeyWhitelist, conf.Server.WhiteList)
	})

	// test router
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "web server ok"})
	})

	// v1
	v1 := router.Group("/v1")
	v1.Use()
	v1.POST("/login")
	v1.POST("/submit")
	v1.POST("/read")

	return router, nil
}
