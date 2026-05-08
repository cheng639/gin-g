package config

import (
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"os"
	"path"
	"time"
)

var logger zerolog.Logger

func Logger() *zerolog.Logger {
	return &logger
}

type LoggerConfig struct {
	Level string `json:"level" yaml:"level"`
}

type LoggerWriter struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	Filename string `json:"filename" yaml:"filename"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`
	// contains filtered or unexported fields
}

func InitLogger() {
	logWriter := &lumberjack.Logger{
		Filename:   path.Join(Config().WorkDir, Config().Server.Name+".log"),
		MaxSize:    Config().LoggerWriter.MaxSize, // megabytes
		MaxBackups: Config().LoggerWriter.MaxBackups,
		MaxAge:     Config().LoggerWriter.MaxAge, // days
	}

	zerolog.TimeFieldFormat = time.DateTime
	//zlog := zerolog.New(logWriter).With().Timestamp().Logger()
	multi := zerolog.MultiLevelWriter(logWriter, os.Stdout)
	zlog := zerolog.New(multi).With().Timestamp().Logger()
	switch Config().Logger.Level {
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

	logger = zlog
}
