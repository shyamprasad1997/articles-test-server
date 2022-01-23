package main

import (
	config "articles-test-server/configs"
	"articles-test-server/infrastructure"
	"articles-test-server/router"
	"articles-test-server/shared/database"
	"articles-test-server/shared/utils"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	logger, err := infrastructure.NewLogger()
	if err != nil {
		panic(err)
	}
	masterDB, readDB, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err, "Failed to initialize DB")
	} else {
		logger.Log.Info("Listen to the API request")
	}
	defer func() {
		rErr := readDB.Close()
		mErr := masterDB.Close()
		if mErr != nil {
			fmt.Println(utils.ErrorsWrap(mErr, "Failed to close master db connection"))
		}
		if rErr != nil {
			fmt.Println(utils.ErrorsWrap(rErr, "Failed to close read db connection"))
		}
	}()

	mux := chi.NewRouter()
	r := &router.Router{
		Mux:           mux,
		MasterDB:      masterDB,
		ReadDB:        readDB,
		LoggerHandler: logger,
	}
	r.InitializeRouter()
	r.SetupHandler()
	server := http.Server{Addr: config.Load().Server.Port, Handler: mux}
	server.SetKeepAlivesEnabled(false)
	server.ListenAndServe()
}
