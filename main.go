package main

import (
	"encoding/json"
	"net/http"
	"safe/users"
)

func signInUser(w http.ResponseWriter, r *http.Request) {
	returningUser := getUser(w, r)
	ok := users.DefaultUserService.VerifyUser(returningUser)

	if !ok {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(returningUser)
	if err != nil {
		return
	}
}

func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(w, r)
	err := users.DefaultUserService.CreateUser(newUser)

	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) users.User {
	var user users.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return users.User{
			Email:    "",
			Password: "",
		}
	}

	return user
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	}
}

func main() {
	http.HandleFunc("/", userHandler)
	err := http.ListenAndServe("", nil)
	if err != nil {
		return
	}
}
