package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kkamdooong/go-restful-api-example/db"
	"github.com/kkamdooong/go-restful-api-example/model"
)

func GetCompanies(w http.ResponseWriter, _ *http.Request) {
	companies := db.FindAll()

	bytes, err := json.Marshal(companies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

// swagger:operation GET /companies/{name} companies getCompany
//
// Get a company by name.
//
// ---
// produces:
// - application/json
//
// schemes:
// - http
//
// parameters:
// - name: name
//   in: path
//   required: true
//   type: string
//
// responses:
//   '200':
//     description: A Company
//     schema:
//       "$ref": "#/responses/Company"
func GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	com, ok := db.FindBy(name)
	if !ok {
		http.NotFound(w, r)
		return
	}

	bytes, err := json.Marshal(com)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

// swagger:operation GET /companies companies getCompanies
//
// Lists companies.
//
// This will show all companies.
// ---
// produces:
// - application/json
//
// schemes:
// - http
//
// responses:
//   '200':
//     description: An array of Companies
//     schema:
//       type: array
//       items:
//         "$ref": "#/responses/Company"
func SaveCompany(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	com := new(model.Company)
	err = json.Unmarshal(body, com)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db.Save(com.Name, com)

	w.Header().Set("Location", r.URL.Path+"/"+com.Name)
	w.WriteHeader(http.StatusCreated)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	com := new(model.Company)
	err = json.Unmarshal(body, com)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db.Save(name, com)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	db.Remove(name)
	w.WriteHeader(http.StatusNoContent)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
