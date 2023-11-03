package main

import "net/http"

func (app *application) getSkinsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get skins"))
}
