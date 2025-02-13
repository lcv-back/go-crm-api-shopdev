package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DBName   string `mapstructure:"dbName"`
	} `mapstructure:"database"`
	Security struct {
		JWT struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read file configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config file: %w", err))
	}

	// read server configuration
	fmt.Println("Server Port:: ", viper.GetInt("server.port"))
	fmt.Println("Jwt key:: ", viper.GetString("security.jwt.key"))

	// configure structure
	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	fmt.Println("Server Port:: ", config.Server.Port)

	// In danh sách các cơ sở dữ liệu và dbName
	for _, db := range config.Databases {
		fmt.Printf(
			"user: %s, password: %s, host: %s, dbName: %s\n", db.User, db.Password, db.Host, db.DBName,
		)
	}
}
