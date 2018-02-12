package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	// "github.com/fenwickelliott/appdir"
	"github.com/fenwickelliott/xplat"
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

type player struct {
	IsPlaying bool `json:"is_playing"`
}

var db, err = xplat.Appdir("snatch")
var user userData

func main() {
	check(err)
	userBytes, err := ioutil.ReadFile(path.Join(db, "userData.json"))
	if os.IsNotExist(err) {
		initialize()
	} else {
		err = json.Unmarshal(userBytes, &user)
		check(err)
	}

	playerBytes := get("me/player")
	var p player
	err = json.Unmarshal(playerBytes, &p)
	check(err)

	if !p.IsPlaying {
		log.Fatal("nothing playing")
	}

	song := getSong()

	isPresant := checkSong(song.ID)
	if isPresant {
		fmt.Println(song.Name, "was already present")
	} else {
		addSong(song)
	}
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

func checkSong(songID string) bool {
	playlist := get("users/" + user.UserID + "/playlists/" + user.PlaylistID + "/tracks")
	return strings.Contains(string(playlist), songID)
}

func addSong(song item) {
	fmt.Println("Adding")
	req, err := http.NewRequest("POST", "https://api.spotify.com/v1/users/"+user.UserID+"/playlists/"+user.PlaylistID+"/tracks?uris=spotify%3Atrack%3A"+song.ID, nil)
	check(err)
	req.Header.Set("Authorization", "Bearer "+user.AcessBearer)
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		fmt.Println(song.Name, "was successfully snatched!")
	}
}
