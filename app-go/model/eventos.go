package model

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"app-go/repository"

	"github.com/gorilla/mux"
)

type Evento struct {
	ID int `json:"id"`
	Titulo string `json:"titulo"`
	Descricao string `json:"descricao"`
	Inicio CustomTime `json:"inicio"`
	Fim CustomTime `json:"fim"`
	CreatedAt CustomTime `json:"created_at"`
	UpdatedAt *CustomTime `json:"updated_at"`
	DeletedAt *CustomTime `json:"deleted_at"`
}

type CustomTime struct {
    time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
    if ct.Time.IsZero() {
        return []byte("null"), nil
    }
    // Formato MySQL: "2006-01-02 15:04:05"
    return json.Marshal(ct.Time.Format("2006-01-02 15:04:05"))
}

var listaDeEventos []Evento

func AdicionaEvento(w http.ResponseWriter, r *http.Request) {
	// Adiciona um evento
}

func ListaEventos(w http.ResponseWriter, r *http.Request) {
	// Lista todos os eventos

	db, err := repository.Connect()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()
	
	rows, err := db.Query("SELECT * FROM eventos")
	if err == sql.ErrNoRows {
        log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
    }
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for rows.Next() {
		var evento Evento
		var updatedAt, deletedAt sql.NullTime
		err = rows.Scan(
						&evento.ID, 
						&evento.Titulo, 
						&evento.Descricao, 
						&evento.Inicio.Time,
						&evento.Fim.Time,
						&evento.CreatedAt.Time,
						&updatedAt,
						&deletedAt,
					)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if updatedAt.Valid {
            customUpdatedAt := CustomTime{updatedAt.Time}
            evento.UpdatedAt = &customUpdatedAt
        }
        if deletedAt.Valid {
            customDeletedAt := CustomTime{deletedAt.Time}
            evento.DeletedAt = &customDeletedAt
        }

        listaDeEventos = append(listaDeEventos, evento)

	}

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listaDeEventos)

	// Limpa a lista de eventos
	listaDeEventos = nil

}	

func DeletarEvento(w http.ResponseWriter, r *http.Request) {	
	// Deleta um evento
	vars := mux.Vars(r)
	id := vars["id"]

	evento, err := umEvento(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if evento == nil {
		err404 := map[string]string{"error": "Evento não encontrado"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err404)
        return
	}
}

func ModificarEvento(w http.ResponseWriter, r *http.Request) {
	// Modifica um evento
	vars := mux.Vars(r)
	id := vars["id"]

	evento, err := umEvento(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if evento == nil {
		err404 := map[string]string{"error": "Evento não encontrado"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err404)
        return
	}

	
}

func RetornaEvento(w http.ResponseWriter, r *http.Request) {
	// Retorna um evento
	vars := mux.Vars(r)
	id := vars["id"]

	evento, err := umEvento(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if evento == nil {
		err404 := map[string]string{"error": "Evento não encontrado"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err404)
        return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(evento)

}

func umEvento(id string) (*Evento, error) {

	db, err := repository.Connect()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	var evento Evento
	var updatedAt, deletedAt sql.NullTime
	err = db.QueryRow("SELECT * FROM eventos WHERE id = ?", id).Scan(
		&evento.ID,
		&evento.Titulo,
		&evento.Descricao,
		&evento.Inicio.Time,
		&evento.Fim.Time,
		&evento.CreatedAt.Time,
		&updatedAt,
		&deletedAt,
	)
	if err == sql.ErrNoRows {
        return nil, nil
    }
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if updatedAt.Valid {
		customUpdatedAt := CustomTime{updatedAt.Time}
		evento.UpdatedAt = &customUpdatedAt
	}

	if deletedAt.Valid {
		customDeletedAt := CustomTime{deletedAt.Time}
		evento.DeletedAt = &customDeletedAt
	}

	return &evento, nil
}