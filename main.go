package main

import (
	"go-challenege/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln("program cannot start: ", err)
	}
}
