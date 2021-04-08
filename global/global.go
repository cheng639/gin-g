package global

import (
	"go.uber.org/zap"

	"gin-g/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	VP     *viper.Viper
	//LOG    *oplogging.Logger
	LOG    *zap.Logger
)
