package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pessoa struct {
	ID        string    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Sobrenome string    `json:"sobrenome,omitempty"`
	Endereco  *Endereco `json:"endereco,omitempty"`
}

type Endereco struct {
	Cidade string `json:"cidade,omitempty"`
	Estado string `json:"estado,omitempty"`
}

var pessoas []Pessoa

func GetPessoas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

func GetPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range pessoas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Pessoa{})
}

func CriarPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var pessoa Pessoa

	_ = json.NewDecoder(r.Body).Decode(&pessoa)
	pessoa.ID = params["id"]
	pessoas = append(pessoas, pessoa)
	json.NewEncoder(w).Encode(pessoas)
}

func ExcluirPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range pessoas {
		if item.ID == params["id"] {
			pessoas = append(pessoas[:index], pessoas[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(pessoas)
}

func main() {
	router := mux.NewRouter()

	pessoas = append(pessoas, Pessoa{ID: "1", Nome: "Alexandre", Sobrenome: "Gutierres"})
	pessoas = append(pessoas, Pessoa{ID: "2", Nome: "Lucia", Sobrenome: "Gutierres"})
	pessoas = append(pessoas, Pessoa{ID: "3", Nome: "Davi", Sobrenome: "Gutierres"})
	pessoas = append(pessoas, Pessoa{ID: "4", Nome: "Danilo", Sobrenome: "Gutierres"})

	router.HandleFunc("/", GetPessoas).Methods("GET")
	router.HandleFunc("/{id}", GetPessoa).Methods("GET")
	router.HandleFunc("/{id}", CriarPessoa).Methods("POST")
	router.HandleFunc("/{id}", ExcluirPessoa).Methods("DELETE")

	erro := http.ListenAndServe(":3000", router)

	if erro != nil {
		log.Fatal(erro)
	}
}
