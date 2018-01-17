package main

import (
	"fmt"
	"os"
	"path"
)

// for Mac, make platform dependant later
var db = path.Join(os.Getenv("HOME"), "Library", "Application Support", "GoSnatch")

func main() {
	fmt.Println(db)
}
