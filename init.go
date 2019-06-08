package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func initialize() {
	fmt.Println("Initializing...")
	done := make(chan bool)
	go serve(done)
	err = Openbrowser("https://accounts.spotify.com/authorize/?client_id=715c15fc7503401fb136d6a79079b50c&response_type=code&redirect_uri=http://localhost:3456/catch&scope=user-read-playback-state%20playlist-read-private%20playlist-modify-private")
	check(err)

	finished := <-done
	getPlaylist()
	if finished {
		fmt.Println("Initiation complete")
	}
}

func serve(done chan bool) {
	http.HandleFunc("/catch", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Thank you, snatch can now access your spotify account.\nYou may close this window.\n")
		code := r.URL.Query()["code"][0]

		getToken(strings.NewReader("grant_type=authorization_code&code=" + code + "&redirect_uri=http://localhost:3456/catch"))

		done <- true
	})
	http.ListenAndServe(":3456", nil)
}

func getPlaylist() {
	listBytes := get("me/playlists")
	list := make(map[string]interface{})
	json.Unmarshal(listBytes, &list)
	items := list["items"].([]interface{})

	for _, v := range items {
		cell := v.(map[string]interface{})
		if cell["name"] == "Snatched" {
			user.PlaylistID = cell["id"].(string)
			owner := cell["owner"].(map[string]interface{})
			user.UserID = owner["id"].(string)
			writeUser()
			return
		}
	}
	createPlaylist()
}

func createPlaylist() {
	meBytes := get("me")
	me := make(map[string]interface{})
	json.Unmarshal(meBytes, &me)
	user.UserID = me["id"].(string)

	url := "https://api.spotify.com/v1/users/" + user.UserID + "/playlists"
	body := strings.NewReader(`{"name":"Snatched","description":"Your automatically generated Snatched playlist!","public":"false"}`)

	req, err := http.NewRequest("POST", url, body)
	check(err)
	req.Header.Set("Authorization", "Bearer "+user.AccessBearer)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	map2b := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &map2b)
	check(err)
	user.PlaylistID = map2b["id"].(string)
	writeUser()
}
