package routes

import (
	"github.com/deyr02/bnzl_crm/pkg/controllers"
	"github.com/gorilla/mux"
)

var Actvity_Category_Routes = func(router *mux.Router) {
	router.HandleFunc("/activity_category/", controllers.GetActivity_Categories).Methods(("GET"))
	router.HandleFunc("/activity_category/", controllers.CreateActivity_Category).Methods(("POST"))
	router.HandleFunc("/activity_category/{activity_categoryId}", controllers.UpdateActivity_Category).Methods("PUT")
	router.HandleFunc("/activity_category/{activity_categoryId}", controllers.GetActivity_CategoryById).Methods("GET")
	router.HandleFunc("/activity_category/{activity_categoryId}", controllers.DeleteActivity_Category).Methods("DELETE")
}
