package main

import (
	"flag"
	"fmt"
	"github.com/Delisa-sama/go-test/news/service"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var appName = "News Service"

func main() {
	flag.Int("port", 80, "Port for web-server")
	flag.String("amqp_url", "amqp://guest:guest@localhost:5672/", "URL string for connect to message broker via amqp")
	flag.Int("storage_timeout", 5, "Timeout in seconds for RPC to storage service")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	fmt.Printf("Starting %v\n", appName)

	service.StartWebServer(viper.GetInt("port"))
}
