package controllers

import (
	"fmt"
	"net/http"
	m "review_utk_uts/models"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateTransactionV1(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "UPDATE transactions SET UserID=?, ProductID=?, Quantity=? WHERE id=?"

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	user_id, errUID := strconv.Atoi(r.Form.Get("user_id"))
	product_id, errPID := strconv.Atoi(r.Form.Get("product_id"))
	quantity, errQty := strconv.Atoi(r.Form.Get("quantity"))

	vars := mux.Vars(r)
	transactionID := vars["transactionID"]
	//transaction_id := r.Form.Get("transaction_id")

	if (errUID != nil || errPID != nil || errQty != nil || transactionID == "") {
		sendErrorResponse(w, 400, "Bad Request! Parameter's missing!")
		return
	}

	result, err := db.Exec(query, user_id, product_id, quantity, transactionID)
	
	if err != nil {
		sendErrorResponse(w, 500, "Internal Server Error! Database Query Fail!")
		return
	}

	amountRowsAffected, _ := result.RowsAffected()

	if (amountRowsAffected == 0) {
		sendErrorResponse(w, 404, "Data not found!")
		return
	}else{
		sendSuccessResponse(w, "Update Transaction Successful!")
	}
}


func UpdateTransactionV2(w http.ResponseWriter, r *http.Request) {
	db := connectGORM()

	if db == nil {
		sendErrorResponse(w, 500, "Internal Server Error! DB Connection Fail!")
		return
	}

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	user_id, errUID := strconv.Atoi(r.Form.Get("user_id"))
	product_id, errPID := strconv.Atoi(r.Form.Get("product_id"))
	quantity, errQty := strconv.Atoi(r.Form.Get("quantity"))

	vars := mux.Vars(r)
	transactionID,_ := strconv.Atoi(vars["transactionID"])

	//ID, errTID := strconv.Atoi(r.Form.Get("transactionID"))

	if (errUID != nil || errPID != nil || errQty != nil) {
		sendErrorResponse(w, 400, "Bad Request! Parameter's missing!")
		return
	}

	transaction := m.Transaction {
		ID: transactionID,
		UserID: user_id,
		ProductID: product_id,
		Quantity: quantity,
	}

	db.First(&transaction, "ID=?", transactionID)
	fmt.Println(transaction.UserID)
	fmt.Println(transaction.ProductID)
	fmt.Println(transaction.Quantity)

	if user_id != 0 {
		transaction.UserID = user_id
	}

	if product_id != 0 {
		transaction.ProductID = product_id
	}

	if quantity != 0 {
		transaction.Quantity = quantity
	}

	result := db.Save(&transaction).Where("ID = ?", transactionID)

	// errQuery := db.Model(&m.Transaction{}).Where("ID = ?", transactionID).Updates(&transaction).Error

	if result.Error != nil {
        sendErrorResponse(w, 500, "Update Transaction Fail")
    } else {
        sendSuccessResponse(w, "Update Transaction Successful")
    }
}

func DeleteTransactionV1(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	vars := mux.Vars(r)
	transactionID := vars["transactionID"]
	//transaction_id := r.Form.Get("transaction_id")

	if (transactionID == "") {
		sendErrorResponse(w, 400, "Bad Request! Parameter's missing!")
		return
	}

	query := "DELETE FROM transactions WHERE ID = ?"

	result, err := db.Exec(query, transactionID)
	if err != nil {
		sendErrorResponse(w, 500, "Delete Transaction Fail")
	} 

	amountRowsAffected, _ := result.RowsAffected()

	if (amountRowsAffected == 0) {
		sendErrorResponse(w, 404, "Data not found!")
		return
	}else{
		sendSuccessResponse(w, "Delete Transaction Successful!")
	}
}

func DeleteTransactionV2(w http.ResponseWriter, r *http.Request) {
	db := connectGORM()

	if db == nil {
		sendErrorResponse(w, 500, "Internal Server Error! DB Connection Fail!")
		return
	}

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	vars := mux.Vars(r)
	transactionID, errTID := strconv.Atoi(vars["transactionID"])

	//ID, errTID := strconv.Atoi(r.Form.Get("transactionID"))

	if (errTID != nil) {
		sendErrorResponse(w, 400, "Bad Request! Parameter's missing!")
		return
	}

	transaction := m.Transaction {
		ID: transactionID,
	}
	
	result := db.Delete(&transaction)

	if result.RowsAffected == 0 {
        sendErrorResponse(w, 404, "Transaction not found")
    } else if result.Error != nil {
        sendErrorResponse(w, 500, "Delete Transaction Fail")
    } else {
        sendSuccessResponse(w, "Delete Transaction Successful")
    }

}

func GetTransactionDetailV2(w http.ResponseWriter, r *http.Request) {
	db := connectGORM()

	if db == nil {
		sendErrorResponse(w, 500, "Internal Server Error! DB Connection Fail!")
		return
	}

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	var transactionsDetail []m.TransactionDetail

	errQuery := db.Table("transactions t").
    Select("t.ID AS transactionID, u.ID AS userID, u.Name AS user_name, u.Age AS user_age, u.Address AS user_address, p.ID AS productID, p.Name AS product_name, p.Price AS product_price, t.quantity").
    Joins("JOIN users u ON t.UserID = u.ID").
    Joins("JOIN products p ON t.ProductID = p.ID").
    Scan(&transactionsDetail).Error

	if errQuery != nil {
		sendErrorResponse(w, 500, "Internal Server Error! Database Query Fail!")
	} else {
		sendSuccessResponse(w, "Success")
	}
}

func GetTransactionDetailV1(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `
		SELECT t.ID AS 'transactionID', u.ID AS 'userID', u.Name AS 'user_name', u.Age AS 'user_age', u.Address AS 'user_address', p.ID AS 'productID', p.Name AS 'product_name', p.Price AS 'product_price', t.quantity
		FROM transactions t
		JOIN users u ON t.UserID = u.ID
		JOIN products p ON t.ProductID = p.ID;`

	transactionDetailRow, err := db.Query(query)
	if err != nil {
		sendErrorResponse(w, 500, "Internal Server Error! Database Query Fail!")
		return
	}

	var transactionDetail m.TransactionDetail
	var transactionDetails []m.TransactionDetail
	for transactionDetailRow.Next(){
		if err := transactionDetailRow.Scan(
		&transactionDetail.ID, &transactionDetail.User.ID, &transactionDetail.User.Name, &transactionDetail.User.Age, &transactionDetail.User.Address, &transactionDetail.Product.ID, &transactionDetail.Product.Name, &transactionDetail.Product.Price, &transactionDetail.Quantity); err != nil {
			print(err.Error())
			return
		} else {
			transactionDetails = append(transactionDetails, transactionDetail)
		}
	}

	if len(transactionDetails) == 0 {
		sendErrorResponse(w, 404, "Data not found")
	}else{
		sendTransactionDetailSuccessResponse(w, "Success", transactionDetails)
	}
}

