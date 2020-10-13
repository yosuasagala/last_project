package user

import (
	"fmt"
	"log"
	"net/http"
)

func (m *Module) LoginUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			log.Println(`error execute template login, err : `, err)
			return
		}

	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" {
			fmt.Println("missing username")
			fmt.Fprintln(w, "missing username!")
		}
		if password == "" {
			fmt.Println("missing password")
			fmt.Fprintln(w, "missing password!")
		}

		_, err := m.Queries.LoginUser.Query(username, password)
		if err != nil {
			log.Println("Failed to insert data")
			return
		}

		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Save(r, w)

		log.Println("berhasil login")

		http.Redirect(w, r, "http://localhost:9090/homelogged", 303)
	}
}
