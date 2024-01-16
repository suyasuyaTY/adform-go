package main

import (
	"adform-go/config"
	"adform-go/server"
	"flag"
)

func main() {
	env := flag.String("e", "development","")
	flag.Parse()

	config.Init(*env)
	if err := server.Init(); err != nil {
		panic(err)
	} 
}