package main

import (
	"ecommerce-ums/cmd"
	"ecommerce-ums/helpers"
)

func main() {
	helpers.SetupConfig()
	helpers.SetupLogger()
	helpers.SetupPostgreSQL()
	// helpers.SetupRedis()
	// cmd.ServerKafkaConsumer()
	cmd.ServeHTTP()
}