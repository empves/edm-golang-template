package fib

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func handleHeartbeat(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func fib(x uint64) uint64 {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func handleFib(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["param"]
	x, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("invalid parameter (expect unsigned integer, got %s)", param)))
		return
	}
	if x > 45 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("parameter too large (expect unsigned integer between 0 and 45, got %d)", x)))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf("%d", fib(x))))
}

func Serve() {
	router := mux.NewRouter()
	router.HandleFunc("/heartbeat", handleHeartbeat).Methods("GET")
	router.HandleFunc("/fib/{param:[0-9]+}", handleFib).Methods("GET")
	http.Handle("/", router)
	port := "8000"
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}
	_ = http.ListenAndServe(":"+port, nil)
}
