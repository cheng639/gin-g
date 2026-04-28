package main

import (
	"errors"
	"gin-g/bootstrap"
	"gin-g/common"
	"gin-g/config"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

func init() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	println("The current working directory is ", workDir)

	config.Config().WorkDir = workDir
	bootstrap.ParseConfig(path.Join(workDir, "config.yaml"))
	bootstrap.InitLogger(config.Logger())

}

func main() {
	defer common.RecoverAndLogPanicStack()

	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		defer common.RecoverAndLogPanicStack()

		panic("panic6")
		panic(errors.New("panic5"))

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
