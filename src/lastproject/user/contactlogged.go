package user

import (
	"fmt"
	"log"
	"net/http"
)

func (m *Module) ContactUserLogged(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "cookie-name")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "contactlogged.html", nil)
		if err != nil {
			log.Println(`error execute template home, err : `, err)
			return
		}
	} else {
		email 	:= r.FormValue("email")
		message := r.FormValue("message")

		if email == "" {
			fmt.Println("missing email")
		}
		if message == "" {
			fmt.Println("missing message")
		}

		_, err := m.Queries.SendMessage.Exec(email, message)
		if err != nil {
			log.Println("Failed to insert data")
		}
		log.Println("berhasil insert data")

		http.Redirect(w, r, "http://localhost:9090/homelogged", 303)
	}
}
