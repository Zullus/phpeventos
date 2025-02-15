package main

import (
	"app-go/model"
	"encoding/json"

	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)


func main(){

	router := mux.NewRouter()

	router.HandleFunc("/src/adiciona-evento", model.AdicionaEvento).Methods("POST")
    router.HandleFunc("/src/lista-eventos", model.ListaEventos).Methods("GET")
    router.HandleFunc("/src/deletar-evento/{id}", model.DeletarEvento).Methods("DELETE")
    router.HandleFunc("/src/modificar-evento/{id}", model.ModificarEvento).Methods("PUT")
    router.HandleFunc("/src/retorna-evento/{id}", model.RetornaEvento).Methods("GET")

	//Retorna erro 404
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Endpoint não encontrado: %s", r.RequestURI)
		err404 := map[string]string{"error": "Endpoint não encontrado"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err404)
	})

	// Configurando o CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),                // Permite todas as origens
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Server running on port 8082")

	log.Fatal(http.ListenAndServe(":8082", corsHandler(router)))
}