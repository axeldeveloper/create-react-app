package person

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Persona struct {
	Name string
	Age  int
}

func Person(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var p Persona

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Do something with the Person struct...
	fmt.Fprintf(w, "Persona: %+v", p)
}
