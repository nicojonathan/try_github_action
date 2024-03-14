package controllers

import (
	"fmt"
	"net/http"
	m "review_utk_uts/models"
	"strconv"
	// "github.com/gorilla/mux"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)


func GetAllUsersV1(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	name := r.Form.Get("name")
	age := r.Form.Get("age")
	address := r.Form.Get("address")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	query := "SELECT * FROM users WHERE 1"

	if name != "" {
		query += fmt.Sprintf(" AND name='%s'", name)
	}

	if age != "" {
		query += fmt.Sprintf(" AND age=%s", age)
	}

	if address != "" {
		query += fmt.Sprintf(" AND address='%s'", address)
	}

	if email != "" {
		query += fmt.Sprintf(" AND email='%s'", email)
	}

	if password != "" {
		query += fmt.Sprintf(" AND password='%s'", password)
	}

	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {		
		sendErrorResponse(w, 500, "Internal Server Error! Database Query Failed!")
		return
	}

	var user m.User
	var users []m.User
	var dataFound bool

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password); err != nil {
			sendErrorResponse(w, 500, "Fail to scan row")
			return
		}else {
			users = append(users, user)
			dataFound = true
		}
	}

	if (!dataFound) {
		sendErrorResponse(w, 404, "Data not found!")
		return
	}

	sendUserSuccessResponse(w, "Get User Successful", users)
}

func GetAllUsersV2(w http.ResponseWriter, r *http.Request) {
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
	age := r.Form.Get("age")
	address := r.Form.Get("address")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	var users []m.User
	query := db.Model(&m.User{})

	if name != "" {
		query = query.Where("name = ?", name)
	}

	if age != "" {
		ageInt,_ := strconv.Atoi(age)
		query = query.Where("age = ?", ageInt)
	}

	if address != "" {
		query = query.Where("address = ?", address)
	}

	if email != "" {
		query = query.Where("email = ?", email)
	}	

	if password != "" {
		query = query.Where("password = ?", password)
	}

	errQuery := query.Find(&users).Error
	if errQuery == nil {

		if len(users) == 0 {
			sendErrorResponse(w, 404, "Data not found")
			return
		}

		sendUserSuccessResponse(w, "Get All User Successful", users)
		return
	}else {
		sendErrorResponse(w, 500, "Get All User Fail")
		return
	}
}
