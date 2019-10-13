package main

import (
	"fmt"
	"github.com/Delisa-sama/go-test/news/service"
)

var appName = "News Service"

func main() {
	fmt.Printf("Starting %v\n", appName)

	service.StartWebServer("80")
}
