package user

import (
	"log"
	"net/http"
)

func (m *Module) AboutUserLogged(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "aboutlogged.html", nil)
		if err != nil {
			log.Println(`error execute template login, err : `, err)
			return
		}

	}
}
