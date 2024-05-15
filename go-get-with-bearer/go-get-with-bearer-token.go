package main

import (
  "fmt"
  "net/http"
)

func main() {

  url := "defin ethe URL here"

  client := &http.Client {
  }
  req, err := http.NewRequest("GET", url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Bearer ******************")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  fmt.Println(res.StatusCode)
}
