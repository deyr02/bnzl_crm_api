package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deyr02/bnzl_crm/pkg/models"
	"github.com/deyr02/bnzl_crm/pkg/utils"
	"github.com/gorilla/mux"
)

var NewActivity_catory models.Activity_Category

func GetActivity_Categories(W http.ResponseWriter, r *http.Request) {
	NewActivity_Categories := models.GetActivity_Categories()
	res, _ := json.Marshal(NewActivity_Categories)
	W.Header().Set("Content-Type", "pkglication/json")
	W.WriteHeader(http.StatusOK)
	W.Write(res)

}

func GetActivity_CategoryById(W http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	activity_categoryId := vars["activity_categoryId"]
	ID, err := strconv.ParseInt(activity_categoryId, 0, 0)

	if err != nil {
		fmt.Println("error whild parsing")
	}
	activity_category_details, _ := models.GetActivity_CategoryById(ID)
	res, _ := json.Marshal(activity_category_details)
	W.Header().Set("Content-Type", "pkglication/json")
	W.WriteHeader(http.StatusOK)
	W.Write(res)
}

func CreateActivity_Category(W http.ResponseWriter, r *http.Request) {

	create_ac := &models.Activity_Category{}
	utils.ParseBody(r, create_ac)
	ac := create_ac.CreateActivity_Category()
	res, _ := json.Marshal(ac)
	W.Header().Set("Content-Type", "pkglication/json")
	W.WriteHeader(http.StatusOK)
	W.Write(res)

}

func DeleteActivity_Category(W http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	activity_categoryId := vars["activity_categoryId"]

	ID, err := strconv.ParseInt(activity_categoryId, 0, 0)

	if err != nil {
		fmt.Println("error whild parsing")
	}
	activity_category_details := models.DeleteActivity_Category(ID)
	res, _ := json.Marshal(activity_category_details)
	W.Header().Set("Content-Type", "pkglication/json")
	W.WriteHeader(http.StatusOK)
	W.Write(res)
}

func UpdateActivity_Category(W http.ResponseWriter, r *http.Request) {

	update_ac := &models.Activity_Category{}
	utils.ParseBody(r, update_ac)
	vars := mux.Vars(r)
	activity_categoryId := vars["activity_categoryId"]

	ID, err := strconv.ParseInt(activity_categoryId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	ac_details, db := models.GetActivity_CategoryById(ID)

	if update_ac.Title != "" {
		ac_details.Title = update_ac.Title
	}

	if update_ac.Sequence != ac_details.Sequence {
		ac_details.Sequence = update_ac.Sequence
	}
	if update_ac.Customizable != ac_details.Customizable {
		ac_details.Customizable = update_ac.Customizable
	}

	db.Save(&ac_details)
	res, _ := json.Marshal(ac_details)
	W.Header().Set("Content-Type", "pkglication/json")
	W.WriteHeader(http.StatusOK)
	W.Write(res)

}
