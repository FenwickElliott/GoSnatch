package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
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
		initialize()
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
		fmt.Println(code)
		done <- true
	})
	http.ListenAndServe(":3456", nil)
}
