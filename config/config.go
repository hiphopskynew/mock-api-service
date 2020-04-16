package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	Config *config
)

type (
	MongoConfig struct {
		URI            string `mapstructure:"uri"`
		DatabaseName   string `mapstructure:"database_name"`
		CollectionName string `mapstructure:"collection_name"`
	}
	config struct {
		MongoConfig `mapstructure:"mongo"`
		ServicePort int `mapstructure:"port"`
	}
)

func LoadConfiguration() {
	vp := viper.New()
	vp.AllowEmptyEnv(true)
	vp.AutomaticEnv()
	vp.AddConfigPath("config")
	vp.SetConfigName("config.yml")
	vp.SetConfigType("yaml")
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	e := vp.ReadInConfig()
	if e != nil {
		panic(fmt.Errorf("Fatal error config file: %s", e.Error()))
	}
	if e := vp.Unmarshal(&Config); e != nil {
		panic(e)
	}
}
