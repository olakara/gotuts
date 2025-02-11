package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {

	// Get environment variables
	os.Getenv("KEY")
	fmt.Println("KEY:", os.Getenv("KEY"))
	os.Setenv("KEY", "VALUE")
	fmt.Println("KEY:", os.Getenv("KEY"))

	viper.SetDefault("ViperKey", "ViperValue")
	fmt.Println("ViperKey:", viper.Get("ViperKey"))
	viper.Set("ViperKey", "ViperValue2")
	fmt.Println("Updated ViperKey:", viper.Get("ViperKey"))

	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file", err)
	}

	fmt.Println("Project from config file:", viper.Get("project"))
	fmt.Println("Version from config file:", viper.Get("version"))
	types := viper.GetStringSlice("types")
	for _, value := range types {
		fmt.Println("Type:", value)
	}
}
