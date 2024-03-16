package main

//Enabling session-based authentication to API endpoints using gorilla/sessions package.

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	redistore "gopkg.in/boj/redistore.v1"
)

var (
	store, _ = redistore.NewRediStore(10, "tcp", ":6379", "", []byte(os.Getenv("session_secret")))
	users    = map[string]string{"naren": "passme", "admin": "password"}
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if originalPassowrd, ok := users[username]; ok {
		session, _ := store.Get(r, "session.id")
		if password == originalPassowrd {
			session.Values["authenticated"] = true
			session.Save(r, w)
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Logged In Successfully"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	session.Options.MaxAge = -1
	session.Save(r, w)
	w.Write([]byte(""))

}

//HealthCheckerHandler returns the date and time.

func HealthCheckerHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	if (session.Values["authenticated"] != nil) && session.Values["authenticated"] != false {
		w.Write([]byte(time.Now().String()))
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/healthcheck", HealthCheckerHandler)
	r.HandleFunc("/logout", LogoutHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		//Good practice enforce timeout for servers you create
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
