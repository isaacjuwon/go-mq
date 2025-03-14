package main

import (
	"log"

	"fusossafuoye.ng/bootstrap"
)


func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(":5001"))
}
