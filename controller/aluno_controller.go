package controller

import (
	"aula4/database"
	"aula4/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerAluno(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postAluno(w, r)

	} else if r.Method == "GET" {
		getAlunos(w, r)
	} else if r.Method == "PUT" {
		putAluno(w, r)
	} else {
		http.Error(w, "Método não implementado", http.StatusNotImplemented)
	}
}

func HandlerAlunoId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//Consulta
		getAlunoId(w, r)

	} else if r.Method == "DELETE" {
		//Exclusao
		deleteAlunoId(w, r)
	}

}

func putAluno(w http.ResponseWriter, r *http.Request) {
	var aluno model.Aluno
	//Qual objeto deve ser atualizado
	err := json.NewDecoder(r.Body).Decode(&aluno)
	
	if err != nil{
		http.Error(w,"requisição inválida", http.StatusBadRequest)
	}

	db, _ := database.GetDatabaseConnection()
	//Save faz update no banco pelo ID de aluno
	db.Save(&aluno)

	fmt.Println(aluno)
	//Devolve como resposta o objeto atualizado
	json.NewEncoder(w).Encode(aluno)
}

func getAlunos(w http.ResponseWriter, r *http.Request) {
	db, _ := database.GetDatabaseConnection()

	var alunos []model.Aluno
	db.Find(&alunos)

	json.NewEncoder(w).Encode(alunos)
}

func deleteAlunoId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db, _ := database.GetDatabaseConnection()
	var aluno model.Aluno

	db.Delete(&aluno, id)

	fmt.Println(&aluno)

	json.NewEncoder(w).Encode(aluno)
}

func getAlunoId(w http.ResponseWriter, r *http.Request) {
	//Capturamos os parametros da requisição
	params := mux.Vars(r)
	//Leitura do parametro {id}
	id := params["id"]
	//Feita a conexão com o db
	db, _ := database.GetDatabaseConnection()
	var aluno model.Aluno
	//Leitura do primeiro objeto que retornar com esse ID
	db.First(&aluno, id)

	fmt.Println(aluno)
	//Transforma o objeto em json para devolver na solicitação
	json.NewEncoder(w).Encode(aluno)
}

func postAluno(w http.ResponseWriter, r *http.Request) {
	//Criamos uma variável/objeto da entidade Aluno
	var aluno model.Aluno

	//Criamos o Decoder para transformar a string json em um objeto(aluno)
	//& = ponteiro , é passado pq o método Decode vai atribuir o valor para o objeto, ou seja alterar o objeto.
	json.NewDecoder(r.Body).Decode(&aluno)

	//Abrimos conexão com o db
	db, err := database.GetDatabaseConnection()
	//Se não foi possivel conectar ao banco retorna erro
	if err != nil {
		http.Error(w, "Erro na conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	//Insere o objeto no banco
	db.Create(&aluno)

	//Escreve no objeto (w)
	json.NewEncoder(w).Encode(aluno)
}
