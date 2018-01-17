package main

import (
	"fmt"
	"os"
	"path"
)

// for Mac, make platform dependant later
var db = path.Join(os.Getenv("HOME"), "Library", "Application Support", "GoSnatch")

type userData struct {
	accessBearer string
	refreshToken string
	userID       string
	playlistID   string
}

func main() {
	user := userData{"access", "refresh", "uID", "pID"}
	fmt.Println(user.userID)
}
