/*
@Time : 2024/12/6 14:20
@Author : linx
@File : ConfigModel.go
@dsc:
*/

package configs

import "gorm.io/gorm"

var Config struct {
	DB       *gorm.DB       `yaml:"-"`
	Logger   loggerConfig   `yaml:"logger"`
	System   systemConfig   `yaml:"system"`
	Database databaseConfig `yaml:"database"`
}

type databaseConfig struct {
	MysqlDataBase `yaml:"mysqlDataBase"`
	RedisDataBase `yaml:"redisDataBase"`
}

type MysqlDataBase struct {
	Host                string `yaml:"host"`
	Port                string `yaml:"port"`
	Username            string `yaml:"username"`
	Password            string `yaml:"password"`
	Database            string `yaml:"db"`
	Charset             string `yaml:"charset"`
	MaxIdleConnections  int    `yaml:"max_idle_connections"`
	MaxOpenConnections  int    `yaml:"max_open_connections"`
	LogMode             string `yaml:"log_mode"`
	EnableFileLogWriter bool   `yaml:"enable_file_log_writer"`
	LogFilename         string `yaml:"log_filename"`
}

type RedisDataBase struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Database    int    `yaml:"db"`
	MaxIdle     int    `yaml:"max_idle"`
	MaxActive   int    `yaml:"max_active"`
	IdleTimeout int    `yaml:"idle_timeout"`
}

type loggerConfig struct {
	LoggerPath     string `yaml:"logger_path"`
	SuccessLogPath string `yaml:"success_path"`
	ErrorLogPath   string `yaml:"error_path"`
	Level          string `yaml:"level"`
	MaxSizes       int    `yaml:"max_sizes"`
	MaxBackups     int    `yaml:"max_backups"`
	MaxAge         int    `yaml:"max_age"`
}

type systemConfig struct {
	Env  string `yaml:"env"`
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}
