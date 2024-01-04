package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func mainTest() {

  url := "http://localhost:8080/videos"
  method := "POST"

  payload := strings.NewReader(`{
    "title": "title2",
    "description": "desc2",
    "url": "url2"
}`)

  client := &http.Client {
  }
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