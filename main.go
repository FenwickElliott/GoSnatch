package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// for Mac, make platform dependant later
var db = path.Join(os.Getenv("HOME"), "Library", "Application Support", "GoSnatch", "userData.json")

type userData struct {
	AcessBearer  string
	RefreshToken string
	UserID       string
	PlaylistID   string
}

func main() {
	userRaw, err := ioutil.ReadFile(db)
	if os.IsNotExist(err) {
		fmt.Println("no such file")
		initialize()
	} else {
		check(err)
	}
	fmt.Println(userRaw)
}

func initialize() {
	user := userData{"Access", "Refresh", "UID", "PID"}
	userJSON, err := json.Marshal(user)
	check(err)
	fmt.Println(string(userJSON))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
