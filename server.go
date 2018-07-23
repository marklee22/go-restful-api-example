// Package main Company API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost:3000
//     BasePath: /api/v1
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kkamdooong/go-restful-api-example/handler"
)

//go:generate swagger generate spec
func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/companies").HandlerFunc(handler.GetCompanies)
	sub.Methods("POST").Path("/companies").HandlerFunc(handler.SaveCompany)
	sub.Methods("GET").Path("/companies/{name}").HandlerFunc(handler.GetCompany)
	sub.Methods("PUT").Path("/companies/{name}").HandlerFunc(handler.UpdateCompany)
	sub.Methods("DELETE").Path("/companies/{name}").HandlerFunc(handler.DeleteCompany)

	log.Fatal(http.ListenAndServe(":3000", router))
}
