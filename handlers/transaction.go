package handlers

import (
	"encoding/json"
	// housedto "housy/dto/house"
	dto "housy/dto/result"
	// usersdto "housy/dto/users"
	// "housy/models"
	"housy/repositories"
	"net/http"
	// "strconv"
	// "github.com/go-playground/validator/v10"
	// "github.com/gorilla/mux"
)

type handlertransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlertransaction {
	return &handlertransaction{TransactionRepository}
}

func (h *handlertransaction) FindTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}
