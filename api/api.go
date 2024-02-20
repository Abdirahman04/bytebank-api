package api

import (
	"fmt"
	"net/http"

	"github.com/Abdirahman04/bytebank-api/api/routes"
	"github.com/Abdirahman04/bytebank-api/logger"
	"github.com/gorilla/mux"
)

func Start() error {
  logger := logger.NewAggregatedLogger()
  logger.Info("Server started")
  router := mux.NewRouter()

  routes.SetUserRoutes(router)
  routes.SetAccountRoutes(router)
  routes.SetTransactionRoutes(router)

  port := ":8080"
  fmt.Printf("Server is running on port %s\n", port)
  err := http.ListenAndServe(port, router)
  if err != nil {
    logger.Error(err.Error())
    return err
  }
  return nil
}
