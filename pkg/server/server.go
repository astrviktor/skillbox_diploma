package server

import (
	"encoding/json"
	"github.com/astrviktor/skillbox_diploma/pkg/result"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const serverAddr = "127.0.0.1:8888"

func listenAndServeHTTP() {
	router := mux.NewRouter()
	router.HandleFunc("/api", handleAPI).Methods("GET")

	log.Fatal(http.ListenAndServe(serverAddr, router))
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	resultT := result.ResultT{Status: false, Error: "Error on collect data"}

	resultSetT := result.GetResultData()
	if result.CheckResult(resultSetT) {
		resultT.Status = true
		resultT.Data = resultSetT
		resultT.Error = ""
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	response, err := json.Marshal(resultT)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func StartServer() {
	listenAndServeHTTP()
}
