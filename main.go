package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// for Mac, make platform dependant later
var db = path.Join(os.Getenv("HOME"), "Library", "Application Support", "GoSnatch", "userData.json")

type userData struct {
	accessBearer string
	refreshToken string
	userID       string
	playlistID   string
}

func main() {
	userRaw, err := ioutil.ReadFile(db)
	if os.IsNotExist(err) {
		fmt.Println("no such file")
	} else {
		check(err)
	}
	fmt.Println(userRaw)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
