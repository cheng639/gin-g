package main

import (
	"errors"
	"fmt"
	"gin-g/bootstrap"
	"gin-g/common"
	"gin-g/config"
	"gin-g/router"
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
	defer common.RecoverAndLogStack()

	engine := gin.New()
	engine.Use(gin.Recovery())
	apiV1 := engine.Group("/api/v1")
	router.RegisterRouters(engine, apiV1.BasePath())

	config.Logger().Info().Msgf("%s is running on %s port.", config.Config().Server.Name, config.Config().Server.Port)
	err := engine.Run(config.Config().Server.Host + ":" + config.Config().Server.Port) // listens on 127.0.0.1:8090 by default
	if err != nil {
		config.Logger().Error().Err(errors.New(fmt.Sprintf("%v", err))).Msgf("%s start failed !", config.Config().Server.Name)
	}
}
