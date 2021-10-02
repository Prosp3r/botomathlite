package main

import "os"

func main() {

	a := App{}

	os.Setenv("BOT_DB_USERNAME", "postgres")
	os.Setenv("BOT_DB_PASWORD", "akomeno123")
	os.Setenv("BOT_DB_NAME", "postgres")

	a.Initialize(
		os.Getenv("BOT_DB_USERNAME"),
		os.Getenv("BOT_DB_PASWORD"),
		os.Getenv("BOT_DB_NAME"))

	a.Run(":8010")
}
