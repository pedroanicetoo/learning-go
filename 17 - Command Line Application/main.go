package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start Point")

	application := app.Generate()

	if erro_r := application.Run(os.Args); erro_r != nil {
		log.Fatal(erro_r)
	}
}
