package api

import (
	"aula4/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InicializeApi() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerHelloWorld)
	r.HandleFunc("/aluno", controller.HandlerAluno)
	r.HandleFunc("/aluno/{id}", controller.HandlerAlunoId).Methods("DELETE", "GET")

	r.HandleFunc("/professor", controller.HadlerProfessor)
	r.HandleFunc("/professor/{id}", controller.HadlerProfessorId).Methods("DELETE", "GET")

	r.HandleFunc("/disciplina", controller.HandleDisciplina)
	r.HandleFunc("/disciplina", controller.HandleDisciplinaId).Methods("DELTE", "GET")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Println(err.Error())
	}

}

func handlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Fprintf(w, "Olá Mundo!")
}

func handlerAluno(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "Em Obras! ! !")

	} else {
		http.Error(w, "Método não implementado", http.StatusNotImplemented)
	}
}
