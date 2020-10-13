package user

import (
	"net/http"
)


func (m *Module) LogoutUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)

	http.Redirect(w, r, "http://localhost:9090/home", 303)
}