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

// type song struct {
// 	Playing bool `json:"is_playing"`
// 	item    struct {
// 		album struct {
// 			Name string `json:"name"`
// 		}
// 	}
// }

type nowPlaying struct {
	Item item `json:"item"`
}

type item struct {
	Album album  `json:"album"`
	Name  string `json:"name"`
	ID    string `json:"id"`
}

type album struct {
	Name string `json:"name"`
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
	getSong()
}

func getSong() {
	var playing nowPlaying
	songBytes := get("me/player/currently-playing")
	// fmt.Println(string(songBytes))

	err := json.Unmarshal(songBytes, &playing)
	check(err)

	fmt.Println(playing)
}
