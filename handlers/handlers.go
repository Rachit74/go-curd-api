package handlers

import (
	"encoding/json"
	"go_curd/models"
	// "go_curd/db"
	// "go_curd/models"
	"net/http"
)

/*

func FunctionName(w http.ResponseWriter, r *http.Request) {
	FunctionBody
}

w -> http.ResponseWrite is use to send response (You can construct an http response with this)
r -> *http.Request is use to recive a request (pointer to the http request struct which contains information about incoming http request)

*/

// getUser Handler function
// This handler gets all the data from the database and writes them to response as json
func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := db.getUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// writes the data as response to json
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.createUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}
