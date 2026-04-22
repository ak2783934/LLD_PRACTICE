2. Implement an idempotent payment handler that takes an idempotency_key and returns the same result for repeated identical requests.



func handleIdempotency(requestID string) (interface{}, error) {
    res, err := requestLog.find(requestID)
    if err == noDataErr {
        return nil, nil 
    }
    if err != nil {
        return nil, err 
    }
    return res, nil 
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
    var err error;
    var resp inteface{}
    string requestID

    ctx, cancel:= context.WithTimeout(context.background(), time.second*5)
    defer cancel()
    defer setIdempotency(resp, requestID, err)

    err := paymentValidation(r)
    if err != nil {
        http.Error(w, "Failed during request validation", http.BadRequest)
        return
    }

    requestID = r.requestID
    idempotencyResp, err := handleIdempotency(r.requestID)
    if err != nil{
        if err == noDataErr {
            continue;
        }
        return err
    }
    if idempotencyResp != nil {
        return idempotencyResp
    }

    resp, err = doPayment(ctx, r)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }

	respMarshal, _ := json.Marshall(resp)
    w.Write(respMarshal)
}