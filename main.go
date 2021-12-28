package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"time"
	"viper-tut/configuration"
)

// main 패키지에 init() 메서드를 만들어놓으면
// main문이 실행되기전에 먼저 실행됩니다.
func init() {
	profile := initProfile()
	setRuntimeConfig(profile)
}

func setRuntimeConfig(profile string) {
	viper.AddConfigPath(".")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&configuration.RuntimeConf)
	if err != nil {
		panic(err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = viper.Unmarshal(&configuration.RuntimeConf)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	viper.WatchConfig()
}

func initProfile() string {
	var profile string
	profile = os.Getenv("GO_PROFILE")
	if len(profile) <= 0 {
		profile = "local"
	}
	fmt.Println("GOLANG_PROFILE: " + profile)
	return profile
}

func main() {
	// 어디서든 가져다 쓸 수 있습니다.
	fmt.Println("db type: ", configuration.RuntimeConf.Datasource.DbType)
	for {
		<-time.After(time.Second * 3)
		fmt.Println("db type: ", configuration.RuntimeConf.Datasource.DbType)
	}
}
