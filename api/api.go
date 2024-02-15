package api

import (
	"fmt"
	"net/http"
  "log"

	"github.com/Abdirahman04/bytebank-api/api/routes"
	"github.com/gorilla/mux"
)

func Start() error {
  router := mux.NewRouter()

  routes.SetUserRoutes(router)
  routes.SetAccountRoutes(router)
  routes.SetTransactionRoutes(router)

  port := ":8080"
  fmt.Printf("Server is running on port %s", port)
  err := http.ListenAndServe(port, router)
  if err != nil {
    log.Fatal("Server failed to start: ", err)
    return err
  }
  return nil
}
