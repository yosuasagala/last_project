package user

import (
	"fmt"
	"log"
	"net/http"
)

type Articles struct {
	ArticleId	int32		`json:"article_id"`
	UserId 		int32		`json:"user_id"`
	Contents    string		`json:"contents"`
	Published	bool		`json:"published"`
}

func (m *Module) HomeUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MASUK 0")
	rows, err := m.Queries.SelectHome.Query()
	if err != nil {
		log.Println("Failed to insert data")
		return
	}
	fmt.Println("MASUK 1")

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

		//log.Println("rowsScanArr :", rowsScanArr)
		//log.Println("rowsScanArr[0] :", rowsScanArr[0])
		//log.Println("rowsScanArr[0].Contents :", rowsScanArr[0].Contents)

	}

	log.Println("berhasil insert data")

	//if r.Method == "GET" {
	err = m.Template.ExecuteTemplate(w, "home.html", rowsScanArr)
	if err != nil {
		log.Println(`error execute template login, err : `, err)
		return
		//	} else {
	}
}
