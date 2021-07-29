package main

import (
	"log"

	"github.com/nikitamirzani323/gofiber_svelte/db"
	"github.com/nikitamirzani323/gofiber_svelte/routes"
)

func main() {
	db.InitMysql()
	app := routes.Init()

	log.Fatal(app.Listen(":4000"))
}
