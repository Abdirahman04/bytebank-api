package main

import (
  "github.com/Abdirahman04/bytebank-api/api"
  "log"
)

func main() {
  err := api.Start()
  if err != nil {
    log.Fatal("Error starting the api", err)
  }
}
