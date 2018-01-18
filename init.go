package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

func initialize() {
	fmt.Println("Initializing...")
	done := make(chan bool)
	go serve(done)
	exec.Command("open", "https://accounts.spotify.com/authorize/?client_id=715c15fc7503401fb136d6a79079b50c&response_type=code&redirect_uri=http://localhost:3456/catch&scope=user-read-playback-state%20playlist-read-private%20playlist-modify-private").Start()

	finished := <-done
	if finished {
		fmt.Println("Initiation complete")
	}
}

func serve(done chan bool) {
	http.HandleFunc("/catch", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Thank you, GoSnatch can now access your spotify account.\nYou may close this window.\n")
		code := r.URL.Query()["code"][0]

		getToken(strings.NewReader("grant_type=authorization_code&code=" + code + "&redirect_uri=http://localhost:3456/catch"))

		done <- true
	})
	http.ListenAndServe(":3456", nil)
}

func getPlaylist() {
	list := get("me/playlists")
	items := list["items"].([]interface{})

	for _, v := range items {
		cell := v.(map[string]interface{})
		if cell["name"] == "GoSnatch" {
			user.PlaylistID = cell["id"].(string)
			owner := cell["owner"].(map[string]interface{})
			user.UserID = owner["id"].(string)
			return
		}
	}
	fmt.Println("TODO: create playlist")
}
