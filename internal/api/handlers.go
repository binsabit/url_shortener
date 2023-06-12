package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app application) ForwardURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.InfoLogger.Println("forward")
}

func (app application) CreateURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.InfoLogger.Println("Create")
}
func (app application) UpdateURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.InfoLogger.Println("Update")
}
func (app application) DeleteURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.InfoLogger.Println("Delete")

}
