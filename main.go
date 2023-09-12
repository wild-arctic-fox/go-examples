package main

import (
	"context"
	"example/server/coinapp"
	"fmt"
	"net/http"
	"time"
)

// controller
func barHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			println("GET")
			client, _ := coinapp.NewHTTPClient(time.Minute)
			client.GetAssets().Data[0].PrintCoinappInfo()
		}
	case http.MethodPost:
		println("POST")
	case http.MethodDelete:
		println("DELETE")
	default:
		println("DEFAULT")
	}
	w.Write([]byte(r.URL.Query().Get("jane")))
}

// Middleware

func contextMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// each req contains own context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", r.URL.Query().Get("id"))

		// rewrite request context
		r = r.WithContext(ctx)

		println(r.Host, r.Method, r.URL.Query().Get("jane"))
		next(w, r)
	}
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		println(r.Host, r.Method, r.URL.Query().Get("jane"), ctx.Value("id").(string))
		next(w, r)
	}
}

func main() {
	fmt.Println("Web Server")

	// handler
	http.HandleFunc("/bar", contextMiddleware(loggerMiddleware(barHandler)))

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()

}
