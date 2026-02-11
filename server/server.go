package server

import (
	"log"
	"net/http"

	"groupieee/handlers"
)

func Start(port string) error {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/artist", handlers.Artist)

	http.HandleFunc("/api/order", handlers.CreateOrder)
	http.HandleFunc("/api/payment", handlers.ProcessPayment)
	http.HandleFunc("/api/payment/paypal", handlers.ProcessPayPalPayment)

	http.HandleFunc("/api/bank-account/setup", handlers.SetupBankAccount)
	http.HandleFunc("/api/bank-account", handlers.GetBankAccount)

	http.HandleFunc("/api/payment/bank-transfer", handlers.ProcessBankTransfer)
	http.HandleFunc("/api/payment/confirm-transfer", handlers.ConfirmBankTransfer)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("HTTP listening on", port)
	return http.ListenAndServe(port, nil)
}
