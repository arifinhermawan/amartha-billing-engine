package router

import (
	"log"
	"net/http"

	"github.com/arifinhermawan/amartha-billing-engine/internal/app/server"
	"github.com/gorilla/mux"
)

func HandleRequest(handlers *server.Handlers) {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(cors)

	handleGetRequest(handlers, router)
	handlePostRequest(handlers, router)

	log.Println("HTTP running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleGetRequest(handlers *server.Handlers, router *mux.Router) {
	// Loan endpoint
	router.HandleFunc("/loans/{loan_id}/outstanding-balance", handlers.Loan.GetOutstandingBalance).Methods("GET")
}

func handlePostRequest(handlers *server.Handlers, router *mux.Router) {
	// Loan endpoint
	router.HandleFunc("/loans", handlers.Loan.CreateLoan).Methods("POST")

	// Payment endpoint
	router.HandleFunc("/payments", handlers.Payment.PayWeeklyInstallment).Methods("POST")

	// User endpoint
	router.HandleFunc("/users", handlers.User.CreateUser).Methods("POST")
}
