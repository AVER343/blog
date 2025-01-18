package main

import (
	"log"
	"net/http"

	"github.com/aver343/blog/pkg/utils"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server Error - %s", r.Method)
	log.Printf("Path - %s", r.URL.Path)
	log.Printf("Err - %s", err)
	utils.WriteJSONError(w, http.StatusInternalServerError, "Internal Server Error , we are working to fix !")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server Error - %s", r.Method)
	log.Printf("Path - %s", r.URL.Path)
	log.Printf("Err - %s", err)
	utils.WriteJSONError(w, http.StatusBadRequest, "Bad Request !")
}
