package main

import (
	"command-line-app/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")
	application := app.Generate()
	if err := application.Run(os.Args); err == nil {
		log.Fatal(err)
	}
}
