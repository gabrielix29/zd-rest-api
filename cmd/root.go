package cmd

import (
	"fmt"
	"github.com/gabrielix29/go-rest-api/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "go-rest-api",
	Short: "RESTful API written in GO",
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config_old.json)")
}

func initConfig() {
	viper.SetConfigType("json")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".") //TODO search better location
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	config.SetDefaults()

	//TODO work with logger
	_ = viper.SafeWriteConfig()
	_ = viper.ReadInConfig()
}
