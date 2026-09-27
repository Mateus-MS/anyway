package main

import "company.anyway/app"

func main() {
	app.GetInstance().Router.Run(":8081")
}
