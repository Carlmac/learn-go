package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("static/css/main.env.css")
	fmt.Println(dir)
	fmt.Println(file)
}
