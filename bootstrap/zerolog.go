package bootstrap

import (
	"gin-g/config"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"os"
	"path"
	"time"
)

func InitLogger(logger *zerolog.Logger) {
	logWriter := &lumberjack.Logger{
		Filename:   path.Join(config.Config().WorkDir, config.Config().Server.Name+".log"),
		MaxSize:    config.Config().LoggerWriter.MaxSize, // megabytes
		MaxBackups: config.Config().LoggerWriter.MaxBackups,
		MaxAge:     config.Config().LoggerWriter.MaxAge, // days
	}

	zerolog.TimeFieldFormat = time.DateTime
	//zlog := zerolog.New(logWriter).With().Timestamp().Logger()
	multi := zerolog.MultiLevelWriter(logWriter, os.Stdout)
	zlog := zerolog.New(multi).With().Timestamp().Logger()
	switch config.Config().Logger.Level {
	case "debug":
		zlog = zlog.Level(zerolog.ErrorLevel)
		break
	case "info":
		zlog = zlog.Level(zerolog.InfoLevel)
		break
	case "warn":
		zlog = zlog.Level(zerolog.WarnLevel)
		break
	case "error":
		zlog = zlog.Level(zerolog.ErrorLevel)
		break
	case "fatal":
		zlog = zlog.Level(zerolog.FatalLevel)
		break
	default:
		zlog = zlog.Level(zerolog.InfoLevel)
	}

	*logger = zlog
}
