package user

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type ArticlesLogin struct {
	ArticleId int32  `json:"article_id"`
	UserId    int32  `json:"user_id"`
	Contents  string `json:"contents"`
	Published bool   `json:"published"`
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func (m *Module) HomeUserLogged(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	rows, err := m.Queries.SelectHome.Query()
	if err != nil {
		log.Println("Failed to insert data")
		return
	}

	var rowsScanArr []Articles

	//Fect data to struct
	for rows.Next() {
		var rowsScan = Articles{}

		err := rows.Scan(&rowsScan.ArticleId, &rowsScan.UserId, &rowsScan.Contents, &rowsScan.Published)
		if err != nil {
			return
		}

		// Append for ervery next row
		rowsScanArr = append(rowsScanArr, rowsScan)
	}

	err = m.Template.ExecuteTemplate(w, "homelogged.html", rowsScanArr)
	if err != nil {
		log.Println(`error execute template login, err : `, err)
		return
	}
}
