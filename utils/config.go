package utils

import "github.com/spf13/viper"

type Environment struct {
	DBName         string `mapstructure:"DB_NAME"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBUser         string `mapstructure:"DB_USER"`
	DBRoot         string `mapstructure:"DB_ROOT"`
	DBRootPassword string `mapstructure:"DB_ROOT_PASSWORD"`
}

func LoadConfig(path string) (env Environment, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}
