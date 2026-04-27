package main

import (
	"fmt"
	"gin-g/common/bootstrap"
	"gin-g/config"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

func init() {
	workDir, err := os.Getwd()
	println("The current working directory is ", workDir)
	if err != nil {
		panic(err)
	}
	config.Config().WorkDir = workDir
	bootstrap.ParseConfig(path.Join(workDir, "config.yaml"))
	println(fmt.Sprintf("Config: %+v", config.Config()))
	
	bootstrap.InitLogger(config.Logger())

}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	config.Logger().Info().Msg("server is running!")
	config.Logger().Error().Msg("server is not running!")
	router.Run() // listens on 0.0.0.0:8080 by default
}
