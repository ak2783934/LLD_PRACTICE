6. Implement a simple HTTP server serving /payments using only the Go standard library, including request validation and basic error handling.

package main

import (
	"fmt"
	"net/http"
)

func paymentHandler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel:= context.WithTimeout(context.background(), time.second*5)
    defer cancel()

    err := paymentValidation(r)
    if err != nil {
        http.Error(w, "Failed during request validation", http.BadRequest)
        return
    }

    resp, err := doPayment(ctx, r)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }

	respMarshal, _ := json.Marshall(resp)
    w.Write(respMarshal)
}

func main() {
	// Register the handler function for the root path "/"
	http.HandleFunc("/payment", paymentHandler)

	// Start the server on port 8080
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
