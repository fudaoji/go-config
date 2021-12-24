package main

import (
	"flag"
	"fmt"

	"github.com/goinggo/mapstructure"
	"github.com/spf13/viper"
)

/**
Get(key string) : interface{}
GetBool(key string) : bool
GetFloat64(key string) : float64
GetInt(key string) : int
GetIntSlice(key string) : []int
GetString(key string) : string
GetStringMap(key string) : map[string]interface{}
GetStringMapString(key string) : map[string]string
GetStringSlice(key string) : []string
GetTime(key string) : time.Time
GetDuration(key string) : time.Duration
IsSet(key string) : bool
AllSettings() : map[string]interface{}
*/

type RedisConf struct {
	db   int
	host string
	port string
	pwd  interface{}
}

func main() {
	mode := flag.String("mode", "prod", "dev mode")
	flag.Parse()
	fmt.Println(*mode)
	viper.SetConfigFile(fmt.Sprintf("./setting_%s.yaml", *mode))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config err: " + err.Error())
	}
	/* var settings = make(map[string]interface{})
	settings = viper.AllSettings() */
	var redis RedisConf
	if err := mapstructure.Decode(viper.Get("redis"), &redis); err != nil {
		fmt.Println(err)
	}
	fmt.Println(viper.Get("redis.port"))
}
