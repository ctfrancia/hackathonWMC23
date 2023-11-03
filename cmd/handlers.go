package main

import "net/http"

func (app *application) getSkinsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get skins"))
}

func (app *application) getMySkinsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get my skins"))
}

func (app *application) getSkinByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get skin by id"))
}

func (app *application) postBuySkinsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buy skin"))
}

func (app *application) putSkinsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("change the skin color"))
}
