package controllers

import (
	"encoding/json"
	"net/http"
	m "review_utk_uts/models"
)

func sendUserSuccessResponse(w http.ResponseWriter, message string, data []m.User) {
    var response m.UsersResponse
    response.Status = 200
    response.Message = message
    response.Data = data

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        // Handle the error appropriately, for example:
        http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        return
    }
}


func sendTransactionDetailSuccessResponse(w http.ResponseWriter, message string, data []m.TransactionDetail) {
    var response m.TransactionDetailResponse
    response.Status = 200
    response.Message = message
    response.Data.Transaction = data

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        // Handle the error appropriately, for example:
        http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        return
    }
}


func sendSuccessResponse(w http.ResponseWriter, message string) {
    var response m.GeneralResponse
    response.Status = 200
    response.Message = message

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        // Handle the error appropriately, for example:
        http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        return
    }
}


func sendErrorResponse(w http.ResponseWriter, status int, message string) {
    var response m.GeneralResponse
    response.Status = status 
    response.Message = message

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        // Handle the error appropriately, for example:
        http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        return
    }
}


