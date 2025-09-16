package api

import (
<<<<<<< HEAD
	"database/sql"  
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/varlaalekya/goproject/model"

	"github.com/IBM/sarama"
)

type Handler struct {
	biz IBizLogic
}

func NewHandler(db *sql.DB, producer sarama.SyncProducer) Handler {
	return Handler{biz: NewBizLogic(db, producer)}
}

func (h Handler) CreateHandler() http.HandlerFunc {
=======
	"database/sql"
	"encoding/json"
	"net/http"

	"apigolang/model"
)

func CreateHandler(db *sql.DB) http.HandlerFunc {
>>>>>>> origin/main
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
<<<<<<< HEAD
		var order model.Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.biz.CreateOrderLogic(order); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h Handler) UpdateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut && r.Method != http.MethodPatch {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var order model.Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// require id for update
		if order.Id == 0 {
			http.Error(w, "id is required", http.StatusBadRequest)
			return
		}
		if err := h.biz.UpdateOrderLogic(order); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h Handler) DeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete && r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// accept id as query (?id=) or JSON body {"id": N}
		idStr := r.URL.Query().Get("id")
		var id int
		if idStr != "" {
			v, err := strconv.Atoi(idStr)
			if err != nil || v <= 0 {
				http.Error(w, "invalid id", http.StatusBadRequest)
				return
			}
			id = v
		} else {
			var body struct {
				Id int `json:"id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			id = body.Id
			if id <= 0 {
				http.Error(w, "invalid id", http.StatusBadRequest)
				return
			}
		}
		if err := h.biz.DeleteOrderLogic(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
=======

		var s model.Student
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := NewBizLogic(db).CreateStudent(s); err != nil {
			http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "student added"})
>>>>>>> origin/main
	}
}
