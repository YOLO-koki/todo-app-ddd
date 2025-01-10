package main

import "todo-app/presentation"

func main() {
	r := presentation.Router()
	r.Run()
}
