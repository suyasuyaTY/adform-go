package main

import (
	"adform-go/server"
)

func main() {
	if err := server.Init(); err != nil {
		panic(err)
	} 
}