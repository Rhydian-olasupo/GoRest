package main

import (
	jsonparse "encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

type JSONServer struct{}

type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	//Read Json file and load data
	absPath, _ := filepath.Abs("books.json")
	raw, readerr := os.ReadFile(absPath)
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}
	//Unmarshall JSON raw data into books array
	masharlerr := jsonparse.Unmarshal(raw, &books)
	if masharlerr != nil {
		log.Println("error:", masharlerr)
		os.Exit(1)
	}

	for _, book := range books {
		if book.ID == args.ID {
			//IF book found, fill reply with it
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	//Create RPC server
	s := rpc.NewServer()
	//Register the Type of data requested as Json
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
