package main

import (
	"example/eve-online-tools/commands"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)
	result, err := commands.ExecuteCommand()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)

}
