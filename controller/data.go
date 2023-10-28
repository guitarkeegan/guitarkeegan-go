package data

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Project struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	P1           string  `json:"p1"`
	P2           string  `json:"p2"`
	P3           string  `json:"p3"`
	Img1         string  `json:"img1"`
	Img2         string  `json:"img2"`
	Img3         string  `json:"img3"`
	Img4         string  `json:"img4"`
	Contributor1 *string `json:"contributor1"`
	Contributor2 *string `json:"contributor2"`
	Contributor3 *string `json:"contributor3"`
	Contributor4 *string `json:"contributor4"`
}

func GetProjectById(projectId int) (Project, error) {

	fp, err := filepath.Abs("mailingList.db")

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", fp)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "select * from project where id = ?"
	row := db.QueryRow(query, projectId)
	var project Project

	err = row.Scan(&project.Id, &project.Title, &project.P1, &project.P2, &project.P3, &project.Img1, &project.Img2, &project.Img3, &project.Img4, &project.Contributor1, &project.Contributor2, &project.Contributor3, &project.Contributor4)

	if err != nil {
		log.Fatal(err)
	}

	return project, nil
}
