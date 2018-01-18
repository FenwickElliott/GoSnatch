package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func get(endpoint string) map[string]interface{} {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/"+endpoint, nil)
	check(err)
	req.Header.Set("Authorization", "Bearer "+user.AcessBearer)
	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	check(err)
	map2b := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &map2b)
	check(err)
	return map2b
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

	userJSON, err := json.Marshal(user)
	check(err)

	err = ioutil.WriteFile(db, userJSON, 0600)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
