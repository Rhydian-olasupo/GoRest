package main

//HANDLING PATH PARAMS AND QUERY PARAMS IN GO
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/*func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])

}*/

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryparams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s!\n", queryparams["id"][0])
	fmt.Fprintf(w, "Got parameter id: %s!\n", queryparams["category"][0])

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", QueryHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
