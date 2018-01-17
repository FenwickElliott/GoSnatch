package main

import (
	"fmt"
	"os"
	"path"
)

var db string

func main() {
	db = path.Join(os.Getenv("HOME"), "Library", "Application Support", "GoSnatch")
	fmt.Println(db)
}
