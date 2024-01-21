package main

//Building a static server file with Golang instead of Using Apache/Nginx
import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	//Mapping to methods is possible with HttpRouter
	router.ServeFiles("/static/*filepath", http.Dir("/Users/user/Desktop/GoRest/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
