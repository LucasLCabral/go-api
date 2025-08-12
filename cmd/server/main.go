package main

import (
	"github.com/LucasLCabral/go-api/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}