package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/api/v1/farms", app.GetFarms)
	router.HandlerFunc(http.MethodPost, "/api/v1/farms", app.InsertFarms)

	router.HandlerFunc(http.MethodPost, "/api/v1/users", app.CreateUser)

	return app.enableCORS(router)
}


