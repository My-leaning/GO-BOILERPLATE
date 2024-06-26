package util

import "github.com/spf13/viper"

type Config struct {
	DBhost     string `mapstructure:"DBHOST"`
	DBport     string `mapstructure:"DBPORT"`
	DBuser     string `mapstructure:"DBUSER"`
	DBpassword string `mapstructure:"DBPASSWORD"`
	DBname     string `mapstructure:"DBNAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
