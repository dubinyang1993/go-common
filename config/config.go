package config

import (
	"github.com/spf13/viper"
	"github.com/lexkong/log"
)

func Init() {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic("Config init failed:" + err.Error())
	}

	//初始化log
	initLog()
}

func initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	if err := log.InitWithConfig(&passLagerCfg); err != nil {
		panic("Log init failed:" + err.Error())
	}
}
