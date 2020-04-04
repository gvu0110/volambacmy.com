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

	http.Redirect(w, r, "/xac-thuc-tai-khoan.html", http.StatusSeeOther)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/adduser", addUser)
	http.ListenAndServe(":8000", nil)
}
