package user

import (
	"fmt"
	"log"
	"net/http"
)

func (m *Module) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "register.html", nil)
		if err != nil {
			log.Println(`error execute template register, err : `, err)
			return
		}
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" {
			fmt.Println("missing username")
		}
		if password == "" {
			fmt.Println("missing department")
		}

		_, err := m.Queries.InsertUser.Exec(username, password)
		if err != nil {
			log.Println("Failed to insert data")
			return
		}

		log.Println("berhasil insert data")

		http.Redirect(w, r, "http://localhost:9090/login", 303)
	}
}
