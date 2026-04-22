1. How does Go’s context package help in a backend service? Show an API handler that uses context for 
timeouts and cancellation.


Ans: 

Its used to pass a lot of information for tracing, and other things into functions in golang. 

lets say if we are passing a trace id, it helps us to trace all the places where that api might have went to. 

In many cases, its also used to pass token, and other non business related things. 

In our case, we pass grab_user_id in many places. 
We also use it for logging. 

So in golang, we use slog.fromCtx(ctx) -> what it does it it prints everything that we have in context. 

It also help us to control the bahaviour of the api in case of timeout and cancellation by the client. 

So when we call an API, we pass some timeout, and when this time is passed, the context is automatically cancelled and api process is further stopped automatically. 

In many cases, we use for cancelling the context by the client, so it kills the api further. 


func paymentHandler(w http.http.ResponseWriter, r *http.Request){

    ctx, cancel := context.WithTimeout(r.context(), 5*time.second)
    defer cancel()

    res, err := doSomeBusinessLogic(ctx, r)
    if err != nil {
        if errors.Is(err, context.Canceled) {
			http.Error(w, "request cancelled", http.StatusRequestTimeout)
			return
		}

        http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

    json.NewEncoder(w).Encode(res)
}


