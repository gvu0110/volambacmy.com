package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var user User

func addUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	f, err := os.OpenFile("addUser.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")

	logger := log.New(f, "", log.LstdFlags)
	logger.Printf("New user: %#v", user)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.FileServer(http.Dir("./static"))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		http.Redirect(w, r, "/error.html", 404)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/adduser", addUser)
	http.ListenAndServe(":8000", nil)
}
