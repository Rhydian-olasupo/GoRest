package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type city struct {
	Name string
	Area uint64
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		fmt.Printf("Got %s city with Area %d sq miles!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - Created"))

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}

func main() {
	http.HandleFunc("/city", postHandler)
	http.ListenAndServe(":8000", nil)
}
