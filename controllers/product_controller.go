package controllers

import (
	"net/http"
	m "review_utk_uts/models"
	"strconv"
	// "github.com/gorilla/mux"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func InsertNewProductV1(w http.ResponseWriter, r *http.Request) {
    db := connect()
	defer db.Close()

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	name := r.Form.Get("name")
	price, errPrice := strconv.Atoi(r.Form.Get("price"))

	if name == "" || errPrice != nil {
		sendErrorResponse(w, 400, "Bad Request! Parameter's missing!")
		return
	}

	query := "INSERT INTO products (`name`, `price`) VALUES (?,?)"

	result, err := db.Exec(query, name, price)
	if err != nil {
		sendErrorResponse(w, 500, "Internal Server Error! Database Query Failed!")
		return
	}

	sendSuccessResponse(w, "Insert Product Successful")
}

func InsertNewProductV2(w http.ResponseWriter, r *http.Request) {
	db := connectGORM()

	if db == nil {
		sendErrorResponse(w, 500, "Internal Server Error! Database connection error")
		return
	}

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	// Read From Query Param
	name := r.Form.Get("name")
	price, errPrice:= strconv.Atoi(r.Form.Get("price"))

	if name == "" || errPrice != nil{
		sendErrorResponse(w, 400, "Bad request! Parameter's missing!")
		return
	}

	product := m.Product{
		Name: name,
		Price: price,
	}
	
	errQuery := db.Select("Name", "Price").Create(&product).Error

	if errQuery == nil {
		sendSuccessResponse(w, "Insert Successful")
		return
	}else {
		sendErrorResponse(w, 500, "Insert Fail")
		return
	}
}