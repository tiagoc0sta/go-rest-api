package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tiagoc0sta/go-rest-api/database"
	"github.com/tiagoc0sta/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)

}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request){
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade) // Recebendo em json e converte para o Go entender (usa NewDecoder)
	database.DB.Create(&novaPersonalidade) 
	json.NewEncoder(w).Encode(novaPersonalidade) // Temos dado json e queremos exibir (usa NewEncoder)
}