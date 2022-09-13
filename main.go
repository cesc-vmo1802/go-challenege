package main

import (
	"go-challenege/cmd"
	"log"
)

// @title Auth Service API
// @version 1.0
// @description This is Auth Service API.

// @contact.name Cesc Nguyen
// @contact.email thuocnv@vmodev.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln("program cannot start: ", err)
	}
}
