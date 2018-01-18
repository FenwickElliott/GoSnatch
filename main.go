package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type userData struct {
	AcessBearer  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string
	PlaylistID   string
}

// for Mac, make platform dependant later
// add functionality for creating GoSnatch directory if not presant
var db = path.Join(os.Getenv("HOME"), "Library", "Application Support", "GoSnatch", "userData.json")
var user userData

func main() {
	userRaw, err := ioutil.ReadFile(db)
	if os.IsNotExist(err) {
		initialize()
	} else {
		check(err)
	}
	err = json.Unmarshal(userRaw, &user)
	check(err)
}
