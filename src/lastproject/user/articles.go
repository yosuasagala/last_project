package user

import (
	"fmt"
	"log"
	"net/http"
)

func (m *Module) ArticlesUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	rows, err := m.Queries.SelectArticles.Query()
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

	err = m.Template.ExecuteTemplate(w, "articles.html", rowsScanArr)
	if err != nil {
		log.Println(`error execute template login, err : `, err)
		return
	}
}

func (m *Module) AddArticlesUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "addarticle.html", nil)
		if err != nil {
			log.Println(`error execute template home, err : `, err)
			return
		}
	} else {
		contents := r.FormValue("contents")

		if contents == "" {
			fmt.Println("missing contents")
		}

		_, err := m.Queries.InsertContents.Exec(contents)
		if err != nil {
			log.Println("Failed to insert data")
			return
		}

		log.Println("berhasil insert data")

		http.Redirect(w, r, "http://localhost:9090/articles", 303)
	}
}

func (m *Module) EditArticlesUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "editarticle.html", nil)
		if err != nil {
			log.Println(`error execute template home, err : `, err)
			return
		}
	} else {
		getId := r.FormValue("id")
		getContent := r.FormValue("contents")

		query, err := m.Queries.EditArticles.Exec(getContent, getId)
		if err != nil {
			log.Println("Failed to edit data")
			return
		}
		fmt.Println(query)
		log.Println("berhasil insert data")

		http.Redirect(w, r, "http://localhost:9090/articles", 303)
	}
}

func (m *Module) RemoveArticlesUser(w http.ResponseWriter, r *http.Request) {
	getId := r.FormValue("id")
	query, err := m.Queries.RemoveArticles.Exec(getId)
	if err != nil {
		log.Println("Failed to delete data", err)
		return
	}

	fmt.Println("QUERYYY: ", query)

	log.Println("berhasil delete data")

	http.Redirect(w, r, "http://localhost:9090/articles", 303)
}

func (m *Module) PublishArticlesUser(w http.ResponseWriter, r *http.Request) {
	getId := r.FormValue("id")

	query, err := m.Queries.PublishArticles.Exec(true, getId)
	if err != nil {
		log.Println("Failed to edit data")
		return
	}
	fmt.Println(query)
	log.Println("berhasil insert data")

	http.Redirect(w, r, "http://localhost:9090/articles", 303)
}

func (m *Module) UnpublishArticlesUser(w http.ResponseWriter, r *http.Request) {
	getId := r.FormValue("id")

	query, err := m.Queries.UnpublishArticles.Exec(false, getId)
	if err != nil {
		log.Println("Failed to edit data")
		return
	}
	fmt.Println(query)
	log.Println("berhasil insert data")

	http.Redirect(w, r, "http://localhost:9090/articles", 303)
}
