package user

import (
	"log"
	"net/http"
)

func (m *Module) AboutUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "about.html", nil)
		if err != nil {
			log.Println(`error execute template login, err : `, err)
			return
		}

	}
}
