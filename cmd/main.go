package main

import (
	"log"

	"github.com/Levap123/test-task-days/cmd/app"
)

func main() {
	app := app.NewApp()
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
	log.Println("server started")
}
