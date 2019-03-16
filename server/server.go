package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/norman/snipit/db"
)

type Snippet struct {
	ID      int    `db:"ID"`
	Content string `json:"content"`
}

type Snippets []*Snippet

func Start(conn *db.Connection) error {
	http.Handle("/snippets", indexHandler(conn))

	if err := http.ListenAndServe(":3000", nil); err != nil {
		return err
	}
	return nil
}

func indexHandler(conn *db.Connection) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rows, err := conn.Session.Queryx(`SELECT * FROM snippets`)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			log.Print(err)
			return
		}
		var snippets Snippets
		for rows.Next() {
			snippet := new(Snippet)
			if err := rows.StructScan(snippet); err != nil {
				http.Error(w, http.StatusText(500), 500)
				log.Print(err)
				return
			}
			snippets = append(snippets, snippet)
		}

		rawJSON, err := json.Marshal(snippets)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			log.Print(err)
			return
		}

		if _, err = w.Write(rawJSON); err != nil {
			http.Error(w, http.StatusText(500), 500)
			log.Print(err)
			return
		}
	})
}
