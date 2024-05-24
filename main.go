package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID        string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World")
	})

	mux.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Listing all the users...")
	})

	mux.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id")
		fmt.Fprintf(w, "Return user with id: %s", id)
	})

	mux.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Post request to /user")
		var u User

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "userData: %+v", u)

		
	})

	//subrouting all to /v1/
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", mux))


	server := http.Server{
		Addr: ":8000",
		Handler: v1,
	}


	fmt.Println("Server Listening on port :8000")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}

}