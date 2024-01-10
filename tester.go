package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func mainTest() {

	url := "http://localhost:8080/videos"
	method := "POST"

	payload := strings.NewReader(`{
    "title": "Cool Title 1",
    "description": "Here is a description for vid 1",
    "url": "https://www.youtube.com/embed/tgbNymZ7vqY",
    "author": {
        "firstname": "Mukhtar",
        "lastname": "Husnain",
        "age": 24,
        "email": "here@email.com"
    }
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic bXVraHRhcjpwYXNza2V5")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
