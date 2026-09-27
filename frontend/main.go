package main

import (
	"company.anyway/app"

	_ "company.anyway/pages/landing"
)

func main() {
	app.GetInstance().Router.Run(":8081")
}
