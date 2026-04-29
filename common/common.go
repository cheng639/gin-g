package common

import (
	"errors"
	"fmt"
	"gin-g/config"
	"log"
	"runtime"
)

func RecoverAndLogStack() {
	if err := recover(); err != nil {
		buf := make([]byte, 1024*1024)
		n := runtime.Stack(buf, false)
		log.Printf("stack: %s", string(buf[:n]))
		config.Logger().Error().Err(errors.New(fmt.Sprintf("%v", err))).Msgf("stack: %s", string(buf[:n]))
	}
}
