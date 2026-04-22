package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"
)

type PaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	UserID   string `json:"user_id"`
}

type PaymentResponse struct {
	Status    string `json:"status"`
	PaymentID string `json:"payment_id"`
}

type IdempotencyRecord struct {
	RequestHash string
	Status      string
	Response    PaymentResponse
	ErrMessage  string
	Done        chan struct{}
}

var (
	mu      sync.Mutex
	records = make(map[string]*IdempotencyRecord)
)

func hashRequest(req PaymentRequest) (string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:]), nil
}

func validatePayment(req PaymentRequest) error {
	if req.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if req.Currency == "" || req.UserID == "" {
		return errors.New("currency and user_id are required")
	}
	return nil
}

func doPayment(ctx context.Context, req PaymentRequest) (PaymentResponse, error) {
	select {
	case <-time.After(200 * time.Millisecond):
		return PaymentResponse{
			Status:    "success",
			PaymentID: "pay_12345",
		}, nil
	case <-ctx.Done():
		return PaymentResponse{}, ctx.Err()
	}
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.Header.Get("Idempotency-Key")
	if key == "" {
		http.Error(w, "missing Idempotency-Key header", http.StatusBadRequest)
		return
	}

	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	if err := validatePayment(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reqHash, err := hashRequest(req)
	if err != nil {
		http.Error(w, "failed to hash request", http.StatusInternalServerError)
		return
	}

	mu.Lock()
	if rec, exists := records[key]; exists {
		if rec.RequestHash != reqHash {
			mu.Unlock()
			http.Error(w, "idempotency key reused with different payload", http.StatusConflict)
			return
		}

		done := rec.Done
		mu.Unlock()

		select {
		case <-done:
			if rec.Status == "SUCCESS" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(rec.Response)
				return
			}
			http.Error(w, rec.ErrMessage, http.StatusInternalServerError)
			return
		case <-time.After(2 * time.Second):
			http.Error(w, "request still processing, retry later", http.StatusAccepted)
			return
		}
	}

	rec := &IdempotencyRecord{
		RequestHash: reqHash,
		Status:      "PROCESSING",
		Done:        make(chan struct{}),
	}
	records[key] = rec
	mu.Unlock()

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := doPayment(ctx, req)

	mu.Lock()
	defer mu.Unlock()

	if err != nil {
		rec.Status = "FAILED"
		rec.ErrMessage = "payment failed"
		close(rec.Done)
		http.Error(w, rec.ErrMessage, http.StatusInternalServerError)
		return
	}

	rec.Status = "SUCCESS"
	rec.Response = resp
	close(rec.Done)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}