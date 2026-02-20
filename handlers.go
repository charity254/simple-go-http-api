package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type User struct {
	ID   int64    `json:"id"`
	Name string  `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type UserStore struct {
	mu sync.RWMutex
	users map[int64] *User
	nextID int64
}
type CreateUserRequest struct {
	Name string `json:"name"`
}

var store = &UserStore{ //REVISIT!!!!!
	users: make(map[int64]*User),
	nextID: 1,
}
func (s *UserStore) Create(name string)*User{
	s.mu.Lock() //prevents concurrent writes
	defer s.mu.Unlock()//unlock even in panic
	u := &User {
		ID : s.nextID,
		Name: name, 
		CreatedAt: time.Now(),
	}
	s.nextID++
	s.users[u.ID]=u
	return u
}
func (s *UserStore) GetById(id int64)(* User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.users[id]
	return u, ok
}
func (s *UserStore) List() []*User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]*User, 0, len(s.users))
	for _, u := range s.users {
		list = append(list, u)
	}
	return list
}
func createUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person CreateUserRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	if person.Name == ""{
		writeError(w, http.StatusBadRequest, "Name is required")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(store.Create(person.Name))
}

func writeError( w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse{
		Error: message,
	})
}
func listUsers(w http.ResponseWriter, r *http.Request){
	
	list := store.List()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}
func getUser(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == ""{
		writeError(w, http.StatusBadRequest, "id is required")
		return
	}
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid id")
		return
	}
	user, ok := store.GetById(Id)
	if !ok {
		writeError(w, http.StatusNotFound, "User Not Found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func getRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	home := "Welcome to the API!\n"
	helloHome := map[string]string{"message": home}
	json.NewEncoder(w).Encode(helloHome)
}

func postGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Only POST method allowed")
		return
	}
	defer r.Body.Close()

	var person User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	if person.Name == ""{
		writeError(w, http.StatusBadRequest, "Name is required")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":"Hello " + person.Name,
	})
}

// func postUsers(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		writeError(w, http.StatusMethodNotAllowed, "Only POST method allowed")
// 		return
// 	}
// 	defer r.Body.Close()

// 	var person User
// 	decoder :=  json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&person); err != nil {
// 		writeError(w, http.StatusBadRequest, "Invalid JSON body")
// 		return
// 	}
// 	if person.Name == ""{
// 		writeError(w, http.StatusBadRequest, "Name is required")
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated) //201 created
// 	json.NewEncoder(w).Encode(person)
// }

func getHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		writeError(w, http.StatusBadRequest, "Name is required")
		return
	}

	greeting := fmt.Sprintf("Hello, %s!\n", name)
	helloData := map[string]string{"message": greeting}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(helloData)
}

func  getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	healthData := map[string]string{"status": "healthy"}
	json.NewEncoder(w).Encode(healthData)
}

var startTime = time.Now()

type StatusResponse struct {
	Service string `json:"service"`
	Uptime string `json:"uptime"`
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime)

	w.Header().Set("Content-Type", "application/json")
	response := StatusResponse {
		Service:"running",
		Uptime: uptime.String(),
	}
	json.NewEncoder(w).Encode(response)
	
}
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost :
			createUser(w, r)
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id == "" {
			listUsers(w, r)
		} else {
			getUser(w, r)
		}
	default:
		writeError(w, http.StatusMethodNotAllowed, "Not A valid Request")
	}
}