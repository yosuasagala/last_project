package user

import (
	"fmt"
	"log"
	"net/http"
)

type Articless struct {
	ArticleId	int32		`json:"article_id"`
	UserId 		int32		`json:"user_id"`
	Contents    string		`json:"contents"`
	Published	bool		`json:"published"`
}

func (m *Module) HomeUsers(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("MASUK 0")
	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "home.html", nil)
		if err != nil {
			log.Println(`error execute template login, err : `, err)
			return
		}
		//else {
		//
		//	rows, err := m.Queries.SelectArticles.Query()
		//	if err != nil {
		//		log.Println("Failed to insert data")
		//		return
		//	}
		//	fmt.Println("MASUK 1")
		//
		//	var rowsScanArr []Articles
		//	//Fect data to struct
		//	for rows.Next() {
		//		var rowsScan = Articles{}
		//
		//		err := rows.Scan(&rowsScan.ArticleId, &rowsScan.UserId, &rowsScan.Contents, &rowsScan.Published)
		//		if err != nil {
		//			return
		//		}
		//
		//		// Append for ervery next row
		//		rowsScanArr = append(rowsScanArr, rowsScan)
		//		log.Println("Contents:", rowsScanArr[0])
		//
		//	}
		//
		//	return
		//
		//
		//
		//	log.Println("berhasil insert data")
		//}
	}
}
