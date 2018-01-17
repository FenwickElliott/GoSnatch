package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

type userData struct {
	AcessBearer  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string
	PlaylistID   string
}

// for Mac, make platform dependant later
var db = path.Join(os.Getenv("HOME"), "Library", "Application Support", "GoSnatch", "userData.json")
var user userData

func main() {
	userRaw, err := ioutil.ReadFile(db)
	if os.IsNotExist(err) {
		initialize()
	} else {
		check(err)
		fmt.Println(userRaw)
	}
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
		ok := exchangeCode(code)
		done <- ok
	})
	http.ListenAndServe(":3456", nil)
}

func exchangeCode(code string) bool {
	body := strings.NewReader("grant_type=authorization_code&code=" + code + "&redirect_uri=http://localhost:3456/catch")
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", body)
	check(err)
	req.Header.Set("Authorization", "Basic NzE1YzE1ZmM3NTAzNDAxZmIxMzZkNmE3OTA3OWI1MGM6ZTkxZWZkZDAzNDVkNDlkNTllOGE2ZDc1YjUzZTE2YTE=")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	check(err)

	err = json.Unmarshal(bodyBytes, &user)
	check(err)
	return true
}
