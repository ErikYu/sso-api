package conf

import "github.com/spf13/viper"

var (
	ServerMode = viper.GetString("server.mode")
	DBPort     = viper.GetString("db_port")
	DBUser     = viper.GetString("db_user")
	DBName     = viper.GetString("db_name")
	DBPassword = viper.GetString("db_password")
)
