package main

import "GoAPIOnECHO/app"

func main() {
	err := app.StartServer()
	if err != nil {
		panic(err)
	}
}
