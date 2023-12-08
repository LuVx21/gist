package config

import (
    "github.com/spf13/viper"
    "os"
)

var Config *viper.Viper

func getConfig() *viper.Viper {
    conf := viper.New()
    conf.SetConfigName("config")
    conf.SetConfigType("yml")
    conf.AddConfigPath("/Users/renxie/OneDrive/Code/gist/Go/gists/config/")
    err := conf.ReadInConfig()
    if err != nil {
        panic(err)
    }
    return conf
}

func Exists(path string) bool {
    _, err := os.Stat(path)
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

func init() {
    Config = getConfig()
}
