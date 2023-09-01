package controller

import (
	"aula4/database"
	"aula4/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleDisciplina(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postDisciplina(w, r)
	} else if r.Method == "GET" {
		getDisciplina(w, r)
	} else if r.Method == "PUT" {
		putDisciplina(w, r)
	}
}

func HandleDisciplinaId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getDisciplinaId(w, r)
	} else if r.Method == "DELETE" {
		deleteDisciplinaId(w, r)
	}
}

func getDisciplinaId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db, _ := database.GetDatabaseConnection()
	var disciplina model.Disciplina

	db.First(&disciplina, id)

	json.NewEncoder(w).Encode(disciplina)
}

func deleteDisciplinaId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db, _ := database.GetDatabaseConnection()
	var disciplina model.Disciplina

	db.Delete(&disciplina, id)

	json.NewEncoder(w).Encode(disciplina)

}

func postDisciplina(w http.ResponseWriter, r *http.Request) {
	var disciplina model.Disciplina

	json.NewDecoder(r.Body).Decode(&disciplina)

	db, err := database.GetDatabaseConnection()

	if err != nil {
		http.Error(w, "Erro na conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	db.Create(&disciplina)

	json.NewEncoder(w).Encode(disciplina)

}

func getDisciplina(w http.ResponseWriter, r *http.Request) {
	db, _ := database.GetDatabaseConnection()
	var disciplina []model.Disciplina

	db.Find(&disciplina)
	json.NewEncoder(w).Encode(disciplina)
}

func putDisciplina(w http.ResponseWriter, r *http.Request) {
	var disciplina model.Disciplina

	err := json.NewDecoder(r.Body).Decode(&disciplina)
	if err != nil {
		http.Error(w, "requisição inválida", http.StatusBadRequest)
	}

	db, _ := database.GetDatabaseConnection()

	db.Save(&disciplina)
	json.NewEncoder(w).Encode(disciplina)

}
