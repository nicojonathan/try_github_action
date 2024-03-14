package main

import (
	"fmt"      // used for formatted printing
	"log"      // used for logging
	"net/http" //  standard library package that provides functionality for building HTTP servers and clients.
	"review_utk_uts/controllers"

	_ "github.com/go-sql-driver/mysql" //  is a blank import to ensure that the MySQL driver is included in the build.
	"github.com/gorilla/mux"           // is used for routing and handling HTTP requests.
)


func main() {
	router := mux.NewRouter()

	// endpoint yang berkaitan dengan user:
	router.HandleFunc("/v1/user", controllers.GetAllUsersV1).Methods("GET")
	router.HandleFunc("/v2/user", controllers.GetAllUsersV2).Methods("GET")

	//endpoint yang berkaitan dengan product:
	router.HandleFunc("/v1/product", controllers.InsertNewProductV1).Methods("POST")
	router.HandleFunc("/v2/product", controllers.InsertNewProductV2).Methods("POST")

	//endpoint yang berkaitan dengan transaction:
	router.HandleFunc("/v1/transaction/{transactionID}", controllers.UpdateTransactionV1).Methods("PUT")
	router.HandleFunc("/v2/transaction/{transactionID}", controllers.UpdateTransactionV2).Methods("PUT")

	router.HandleFunc("/v1/transaction/{transactionID}", controllers.DeleteTransactionV1).Methods("DELETE")
	router.HandleFunc("/v2/transaction/{transactionID}", controllers.DeleteTransactionV2).Methods("DELETE")

	router.HandleFunc("/v2/transaction/detail", controllers.GetTransactionDetailV1).Methods("GET")

	http.Handle("/", router)

	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	//http.ListenAndServe(":8888", router)

	// Start the HTTP server
    server := &http.Server{
        Addr:    ":8888",
        Handler: router,
    }

    fmt.Println("Starting server on port 8888")
    log.Println("Starting server on port 8888")
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("Error starting server: %v", err)
    }

}