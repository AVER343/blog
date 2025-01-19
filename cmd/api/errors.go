package main

import (
	"net/http"

	"github.com/aver343/blog/pkg/utils"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.Logger.Infof("Internal Server Error - %s", r.Method)
	app.Logger.Infof("Path - %s", r.URL.Path)
	app.Logger.Infof("Err - %s", err)
	utils.WriteJSONError(w, http.StatusInternalServerError, "Internal Server Error , we are working to fix !")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	app.Logger.Infof("Internal Server Error - %s", r.Method)
	app.Logger.Infof("Path - %s", r.URL.Path)
	app.Logger.Infof("Err - %s", err)
	utils.WriteJSONError(w, http.StatusBadRequest, err.Error())
}
