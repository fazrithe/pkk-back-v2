package routes

import (
	controllers "pkk-back-v2/app/controllers"
	"pkk-back-v2/app/controllers/auth"

	"github.com/gorilla/mux"
)

func ApiRoutes(prefix string, r *mux.Router) {

	s := r.PathPrefix(prefix).Subrouter()

	s.HandleFunc("/login", auth.Login).Methods("POST")
	s.HandleFunc("/register", controllers.CreateUser).Methods("POST")

	//users
	s.HandleFunc("/users", auth.ValidateMiddleware(controllers.GetUsers)).Methods("GET")
	s.HandleFunc("/users/{id}", auth.ValidateMiddleware(controllers.GetUser)).Methods("GET")
	s.HandleFunc("/users", auth.ValidateMiddleware(controllers.CreateUser)).Methods("POST")
	s.HandleFunc("/users/{id:[0-9]+}", auth.ValidateMiddleware(controllers.GetUser)).Methods("GET")
	s.HandleFunc("/users/{id:[0-9]+}", auth.ValidateMiddleware(controllers.UpdateUser)).Methods("PUT")
	s.HandleFunc("/users/{id}", auth.ValidateMiddleware(controllers.DeleteUser)).Methods("DELETE")

	//institutions
	s.HandleFunc("/institutions", auth.ValidateMiddleware(controllers.GetInstitutions)).Methods("GET")
	s.HandleFunc("/institutions", auth.ValidateMiddleware(controllers.CreateInstitutions)).Methods("POST")
	s.HandleFunc("/institutions/{id}", auth.ValidateMiddleware(controllers.GetInstitution)).Methods("GET")
	s.HandleFunc("/institutions/{id:[0-9]+}", auth.ValidateMiddleware(controllers.GetInstitution)).Methods("GET")
	s.HandleFunc("/institutions/{id:[0-9]+}", auth.ValidateMiddleware(controllers.UpdateInstitution)).Methods("PUT")
	s.HandleFunc("/institutions/{id}", auth.ValidateMiddleware(controllers.DeleteInstitution)).Methods("DELETE")
}
