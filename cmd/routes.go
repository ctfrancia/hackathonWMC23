package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/skins/available", app.getSkinsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/skins/myskins", app.getMySkinsHandler)
	route.HandlerFunc(http.MethodGet, "/v1/skin/getskin/:id", app.getSkinByIdHandler)

	router.HandlerFunc(http.MethodPost, "/v1/skins/buy", app.postBuySkinsHandler)

	router.HandlerFunc(http.MethodPut, "/v1/skins/color", app.putSkinsHandler)

	router.HandlerFunc(http.MethodDelete, "/v1/skins/delete/:id", app.deleteSkinsHandler)

	return router
}
