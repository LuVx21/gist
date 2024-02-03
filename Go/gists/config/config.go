package config

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
    "os"
)

type Config struct {
    Server ServerConfig
    Log    LogConfig
    MySQL  MySQL
    Redis  Redis
}

var AppConfig Config

func config() Config {
    //viper := viper.New()
    viper.SetConfigName("config")
    viper.SetConfigType("yml")
    viper.AddConfigPath("/Users/renxie/OneDrive/Code/gist/Go/gists/config/")
    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }

    viper.OnConfigChange(func(e fsnotify.Event) {
        fmt.Println("Config file changed:", e.Name)
    })
    viper.WatchConfig()

    var c Config
    viper.Unmarshal(&c)
    return c
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
    AppConfig = config()
}
