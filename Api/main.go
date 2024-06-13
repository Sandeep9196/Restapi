package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	users = map[string]User{}
)

func main() {
	http.HandleFunc("/createUser", addUser)
	http.HandleFunc("/getUser", getUser)

	fmt.Println("Users:", users)
	fmt.Println("Server Started on port 8000")

	log.Fatalf("Server failed to start, error: %v", http.ListenAndServe(":8000", nil))
}

// create /add new record of user in map
func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users[user.Name] = user

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Users:", users)
}

// retrieve a user from the map
func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	user, exists := users[name]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
