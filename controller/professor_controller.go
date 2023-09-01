package controller

import (
	"aula4/database"
	"aula4/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func HadlerProfessor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//Incluir
		postProfessor(w, r)
	} else if r.Method == "PUT" {
		//Atualiza
		putProfessor(w, r)
	} else if r.Method == "GET" {
		getProfessor(w, r)
	}
}
func HadlerProfessorId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//Consulta
		getProfessorId(w, r)
	} else if r.Method == "DELETE" {
		//Deletar
		deleteProfessorId(w, r)
	}
}

func postProfessor(w http.ResponseWriter, r *http.Request) {
	var professor model.Professor

	json.NewDecoder(r.Body).Decode(&professor)

	db, _ := database.GetDatabaseConnection()

	db.Create(&professor)

	json.NewEncoder(w).Encode(professor)
}

func getProfessorId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db, _ := database.GetDatabaseConnection()

	var professor model.Professor

	db.First(&professor, id)
	json.NewEncoder(w).Encode(professor)
}

func deleteProfessorId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db, _ := database.GetDatabaseConnection()

	var professor model.Professor

	db.Delete(&professor, id)

	json.NewEncoder(w).Encode(professor)

}

func putProfessor(w http.ResponseWriter, r *http.Request) {
	var professor model.Professor

	err := json.NewDecoder(r.Body).Decode(&professor)

	if err != nil {
		http.Error(w, "requisição inválida", http.StatusBadRequest)
	}

	db, _ := database.GetDatabaseConnection()

	db.Save(&professor)

	json.NewEncoder(w).Encode(professor)
}

func getProfessor(w http.ResponseWriter, r *http.Request) {
	db, _ := database.GetDatabaseConnection()
	var professor []model.Professor

	db.Find(&professor)

	json.NewEncoder(w).Encode(professor)
}
