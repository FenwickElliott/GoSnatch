package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

// func exchangeCode(code string) bool {
// 	fmt.Println("exchangeing")
// 	body := strings.NewReader("grant_type=authorization_code&code=" + code + "&redirect_uri=http://localhost:3456/catch")
// 	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", body)
// 	check(err)
// 	req.Header.Set("Authorization", "Basic NzE1YzE1ZmM3NTAzNDAxZmIxMzZkNmE3OTA3OWI1MGM6ZTkxZWZkZDAzNDVkNDlkNTllOGE2ZDc1YjUzZTE2YTE=")
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	resp, err := http.DefaultClient.Do(req)
// 	check(err)
// 	defer resp.Body.Close()

// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	check(err)

// 	err = json.Unmarshal(bodyBytes, &user)
// 	check(err)

// 	// userJSON, err := json.Marshal(user)
// 	// check(err)

// 	// err = ioutil.WriteFile(db, userJSON, 0600)
// 	// check(err)
// 	return true
// }

// func refresh() {
// 	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader("grant_type=refresh_token&refresh_token="+user.RefreshToken))
// 	check(err)
// 	req.Header.Set("Authorization", "Basic NzE1YzE1ZmM3NTAzNDAxZmIxMzZkNmE3OTA3OWI1MGM6ZTkxZWZkZDAzNDVkNDlkNTllOGE2ZDc1YjUzZTE2YTE=")
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	resp, err := http.DefaultClient.Do(req)
// 	check(err)
// 	defer resp.Body.Close()

// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	check(err)

// 	err = json.Unmarshal(bodyBytes, &user)
// 	check(err)

// 	// userJSON, err := json.Marshal(user)
// 	// check(err)

// 	// err = ioutil.WriteFile(db, userJSON, 0600)
// 	// check(err)
// }

func getToken(body *strings.Reader) {
	fmt.Println("getting")
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

	userJSON, err := json.Marshal(user)
	check(err)

	err = ioutil.WriteFile(db, userJSON, 0600)
	check(err)
}
