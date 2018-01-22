package main

import (
	"encoding/json"
	"fmt"
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

type nowPlaying struct {
	Item item `json:"item"`
}

type item struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// for Mac, make platform dependant later
// add functionality for creating GoSnatch directory if not presant
var db = path.Join(os.Getenv("HOME"), "Library", "Application Support", "GoSnatch", "userData.json")
var user userData

func main() {
	userBytes, err := ioutil.ReadFile(db)
	if os.IsNotExist(err) {
		initialize()
	} else {
		err = json.Unmarshal(userBytes, &user)
		check(err)
	}
	song := getSong()

	fmt.Println(song)
}

func getSong() item {
	var playing nowPlaying
	songBytes := get("me/player/currently-playing")

	err := json.Unmarshal(songBytes, &playing)
	check(err)

	if playing.Item.ID == "" {
		fmt.Println("Nothing playing")
		os.Exit(0)
	}
	return playing.Item
}
