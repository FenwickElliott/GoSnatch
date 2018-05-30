package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func get(endpoint string) []byte {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/"+endpoint, nil)
	check(err)
	req.Header.Set("Authorization", "Bearer "+user.AccessBearer)
	resp, err := http.DefaultClient.Do(req)
	if resp == nil {
		fmt.Println("Internet connection unavailable")
		os.Exit(0)
	}
	check(err)
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		getToken(strings.NewReader("grant_type=refresh_token&refresh_token=" + user.RefreshToken))
		main()
		os.Exit(0)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	check(err)
	return bodyBytes
}

func getToken(body *strings.Reader) {
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

	writeUser()
}

func writeUser() {
	if _, err := os.Stat(db); os.IsNotExist(err) {
		err = os.MkdirAll(db, 0764)
		check(err)
	}

	userJSON, err := json.Marshal(user)
	check(err)

	err = ioutil.WriteFile(path.Join(db, "userData.json"), userJSON, 0600)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
