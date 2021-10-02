package main

import (
	"os"
)

func init() {
	os.Setenv("BOT_DB_USERNAME", "postgres")
	os.Setenv("BOT_DB_PASWORD", "akomeno123")
	os.Setenv("BOT_DB_NAME", "postgres")
}

func main() {

	// a := app.App{}

	// a.Initialize(
	// 	os.Getenv("BOT_DB_USERNAME"),
	// 	os.Getenv("BOT_DB_PASWORD"),
	// 	os.Getenv("BOT_DB_NAME"))

	// a.Run(":8080")
}
