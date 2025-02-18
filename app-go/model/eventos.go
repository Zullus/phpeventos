package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"app-go/repository"

	"github.com/gorilla/mux"
)

type Evento struct {
	ID        int         `json:"id"`
	Titulo    string      `json:"titulo"`
	Descricao string      `json:"descricao"`
	Inicio    CustomTime  `json:"inicio"`
	Fim       CustomTime  `json:"fim"`
	CreatedAt CustomTime  `json:"created_at"`
	UpdatedAt *CustomTime `json:"updated_at"`
	DeletedAt *CustomTime `json:"deleted_at"`
}

type EventoNovo struct {
	ID        string      `json:"id"`
	Titulo    string      `json:"titulo"`
	Descricao string      `json:"descricao"`
	Inicio    CustomTime  `json:"inicio"`
	Fim       CustomTime  `json:"fim"`
}

type CustomTime struct {
	time.Time
}

func (ct CustomTime) Value() (driver.Value, error) {
    return ct.Time, nil
}

func (ct *CustomTime) Scan(value interface{}) error {
    if value == nil {
        return nil
    }
    
    switch v := value.(type) {
    case time.Time:
        ct.Time = v
        return nil
    case string:
        t, err := time.Parse(time.RFC3339, v)
        if err != nil {
            return err
        }
        ct.Time = t
        return nil
    }
    return fmt.Errorf("cannot scan type %T into CustomTime", value)
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	// Formato MySQL: "2006-01-02 15:04:05"
	return json.Marshal(ct.Time.Format("2006-01-02 15:04:05"))
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
    s := strings.Trim(string(b), "\"")
    t, err := time.Parse("2006-01-02T15:04", s)
    if err != nil {
        return err
    }
    ct.Time = t
    return nil
}

var listaDeEventos []Evento

func AdicionaEvento(w http.ResponseWriter, r *http.Request) {
	// Adiciona um evento

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var evento Evento
    err = json.Unmarshal(body, &evento)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	db, err := repository.Connect()
	if err != nil {
		log.Panicln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO eventos (titulo, descricao, inicio, fim) VALUES (?, ?, ?, ?)", evento.Titulo, evento.Descricao, evento.Inicio.Time, evento.Fim.Time)
	if err != nil {
		log.Panicln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Evento adicionado com sucesso")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(evento)

}

func ListaEventos(w http.ResponseWriter, r *http.Request) {
	// Lista todos os eventos

	db, err := repository.Connect()
	if err != nil {
		log.Panicln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM eventos WHERE deleted_at IS NULL")
	if err == sql.ErrNoRows {
		log.Panicln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err != nil {
		log.Panicln(err)
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
			log.Panicln(err)
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
		log.Panicln(err)
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

	if id == "" {
		err404 := map[string]string{"error": "O id do evento é obrigatório"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err404)
		return
	}

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

	db, err := repository.Connect()
	if err != nil {
		log.Panicln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE eventos SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		log.Panicln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	eventoApagado := map[string]string{
		"sucess":  "true",
		"message": "Evento deletado com sucesso"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eventoApagado)

}

func ModificarEvento(w http.ResponseWriter, r *http.Request) {
	// Modifica um evento

    // Ler o corpo da requisição como uma string
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Imprimir o corpo da requisição no terminal
    log.Println("Corpo da requisição:\n", string(body))

	var eventoNovo EventoNovo
    err = json.Unmarshal(body, &eventoNovo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	evento, err := umEvento(eventoNovo.ID)
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

	db, err := repository.Connect()
	if err != nil {
		log.Panicln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE eventos SET titulo= ?,descricao= ?,inicio= ?,fim= ? WHERE id = ?", 
							eventoNovo.Titulo, 
							eventoNovo.Descricao,
							eventoNovo.Inicio,
							eventoNovo.Fim,
							evento.ID)
	if err != nil {
		log.Panicln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	eventoAlterado := map[string]string{
		"id":  string(evento.ID),
		"message": "Evento alterado com sucesso"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eventoAlterado)

}

func RetornaEvento(w http.ResponseWriter, r *http.Request) {
	// Retorna um evento
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		err404 := map[string]string{"error": "O id do evento é obrigatório"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err404)
		return
	}

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
		log.Panicln(err)
		return nil, err
	}

	defer db.Close()

	var evento Evento
	var updatedAt, deletedAt sql.NullTime
	err = db.QueryRow("SELECT * FROM eventos WHERE id = ? AND deleted_at IS NULL", id).Scan(
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
		log.Panicln(err)
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
