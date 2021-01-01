package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Person struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

const (
	host     = "postgres-srv"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GETInitDb(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := OpenConnection()

	sqlStatement := `create table person(id serial primary key, name varchar(40), nickname varchar(40));`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()

}

func GETAllPersons(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}

	var people []Person

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.Name, &person.Nickname)
		people = append(people, person)
	}

	peopleBytes, _ := json.MarshalIndent(people, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer rows.Close()
	defer db.Close()
}

func GETOnePerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := OpenConnection()

	sqlSelect := "SELECT * FROM person WHERE id = $1;"
	row := db.QueryRow(sqlSelect, ps.ByName("id"))

	var person Person

	err := row.Scan(&person.Id, &person.Name, &person.Nickname)

	if err != nil {
		panic(err)
	}

	peopleBytes, _ := json.MarshalIndent(person, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer db.Close()
}

func POSTHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := OpenConnection()
	defer db.Close()

	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO person (name, nickname) VALUES ($1, $2) returning id, name, nickname;`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var person Person;
	err = stmt.QueryRow(p.Name, p.Nickname).Scan(&person.Id, &person.Name, &person.Nickname)
	if err != nil {
		panic(err)
	}

	peopleBytes, _ := json.MarshalIndent(person, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)


}

func main() {
	router := httprouter.New()

	router.GET("/init", GETInitDb)
	router.GET("/person/:id", GETOnePerson)
	router.GET("/person", GETAllPersons)
	router.POST("/person", POSTHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
