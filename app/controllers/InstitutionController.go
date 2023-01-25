package controllers

import (
	"encoding/json"
	"net/http"
	"pkk-back-v2/app/models"
	u "pkk-back-v2/utils"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//get all the institutions
func GetInstitutions(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "success")
	institutions := models.GetInstitutions()
	if institutions == nil {
		u.Respond(w, u.Message(false, "No Institutions found"))
		return
	}
	resp["data"] = institutions
	u.Respond(w, resp)
	return
}

//create institution
func CreateInstitutions(w http.ResponseWriter, r *http.Request) {
	institution := &models.Institution{}

	err := json.NewDecoder(r.Body).Decode(institution)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
	}

	defer r.Body.Close()

	resp := institution.Create()
	u.Respond(w, resp)
}

func GetInstitution(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in yout request"))
		return
	}

	institution := models.GetInstitution(id)
	if institution == nil {
		u.Respond(w, u.Message(false, "User not found"))
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = institution
	u.Respond(w, resp)
	return
}

func UpdateInstitution(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var institution models.Institution
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
	}

	err = models.GetInstitutionForUpdateOrDelete(id, &institution)
	if err != nil {
		u.Respond(w, u.Message(false, "user not found"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&institution)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	institution.ID = uint(id)
	institution.UpdatedAt = time.Now().Local()
	defer r.Body.Close()

	err = models.UpdateInstitution(&institution)
	if err != nil {
		u.Respond(w, u.Message(false, "Could not update the record"))
	}

	resp := u.Message(true, "Update successfully")
	resp["data"] = institution
	u.Respond(w, resp)

}

func DeleteInstitution(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var institution models.Institution
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	err = models.GetInstitutionForUpdateOrDelete(id, &institution)
	if err != nil {
		u.Respond(w, u.Message(false, "User not found"))
		return
	}

	err = models.DeleteInstitution(&institution)
	if err != nil {
		u.Respond(w, u.Message(false, "Could not delete the record"))
		return
	}
	u.Respond(w, u.Message(true, "User has been deleted successfully"))
	return
}
