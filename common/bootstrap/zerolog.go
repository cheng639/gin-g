package bootstrap

import (
	"gin-g/config"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"path"
)

func InitLogger(logger *zerolog.Logger) {
	logWriter := &lumberjack.Logger{
		Filename:   path.Join(config.Config().WorkDir, config.Config().Server.Name+".log"),
		MaxSize:    config.Config().LoggerWriter.MaxSize, // megabytes
		MaxBackups: config.Config().LoggerWriter.MaxBackups,
		MaxAge:     config.Config().LoggerWriter.MaxAge, // days
	}
	zlog := zerolog.New(logWriter).With().Timestamp().Logger()
	zlog = zlog.Level(zerolog.ErrorLevel)
	*logger = zlog
}
