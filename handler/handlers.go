package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kkamdooong/go-restful-api-example/db"
	"github.com/kkamdooong/go-restful-api-example/model"
)

// swagger:operation GET /companies companies getCompanies
//
// Lists companies.
//
// This will show all companies.
// ---
// produces:
// - application/json
//
// responses:
//   200:
//     description: An array of Companies
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Company"
func GetCompanies(w http.ResponseWriter, _ *http.Request) {
	companies := db.FindAll()

	bytes, err := json.Marshal(companies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

// swagger:operation GET /companies/{name} companies getCompanyByName
//
// Get a company by name.
// Gets the details for a company.
//
// ---
// produces:
// - application/json
//
// parameters:
// - name: name
//   in: path
//   required: true
//   type: string
//
// responses:
//   200:
//     description: A Company
//     schema:
//       "$ref": "#/definitions/Company"
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

// swagger:operation POST /companies companies createCompany
//
// Creates a company.
//
// ---
// consumes:
// - application/x-www-form-urlencoded
//
// produces:
// - application/json
//
// parameters:
// - name: name
//   description: Name of the company
//   in: body
//   required: true
//   type: string
// - name: tel
//   description: Telephone of the company
//   in: body
//   required: true
//   type: string
// - name: email
//   description: Email of the company
//   in: body
//   required: true
//   type: string
//
// responses:
//   201:
//     description: Company created
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
}

// swagger:operation PUT /companies/{name} companies updateCompany
//
// Updates a company.
//
// ---
// consumes:
// - application/x-www-form-urlencoded
//
// produces:
// - application/json
//
// parameters:
// - name: name
//   description: Name of the company
//   in: path
//   required: true
//   type: string
// - name: tel
//   description: Telephone of the company
//   in: body
//   required: false
//   type: string
// - name: email
//   description: Email of the company
//   in: body
//   required: false
//   type: string
//
// responses:
//   200:
//     description: Company updated
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

// swagger:operation DELETE /companies/{name} companies deleteCompany
//
// Deletes a company.
//
// ---
// produces:
// - application/json
//
// parameters:
// - name: name
//   description: Name of the company
//   in: path
//   required: true
//   type: string
//
// responses:
//   204:
//     description: Company deleted
func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	db.Remove(name)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNoContent)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
}
