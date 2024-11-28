package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress      string `mapstructure:"server-address"`
	ContractAddress    string `mapstructure:"contract-address"`
	OwnerPublicKey     string `mapstructure:"owner-public-key"`
	OwnerPrivateKey    string `mapstructure:"owner-private-key"`
	TestUserPublicKey  string `mapstructure:"test_user_public_key"`
	TestUserPirvateKey string `mapstructure:"test_user_pirvate_key"`
}

type Using struct {
	Chain       string `mapstructure:"chain"`
	Environment string `mapstructure:"environment"`
}

// var GlobalConfig ChainConfig
var config *Config

func init() {
	v := viper.New()
	v.SetConfigFile("config.yaml")

	// Read configuration file
	err := v.ReadInConfig()
	if err != nil {
		// Handle configuration file reading error
		panic(err)
	}

	// Read Using data
	var using Using
	err = v.UnmarshalKey("using", &using)
	if err != nil {
		// Handle Using configuration parsing error
		panic(err)
	}

	// Build the key for EnvironmentConfig field
	configKey := fmt.Sprintf("%s.%s", using.Chain, using.Environment)

	// Parse the corresponding field into EnvironmentConfig struct type
	err = v.UnmarshalKey(configKey, &config)
	if err != nil {
		// Handle EnvironmentConfig configuration parsing error
		panic(err)
	}
	if config == nil {
		panic("config is nil")
	}
}
