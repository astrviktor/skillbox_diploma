package simulator

import (
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const DataDir = "/home/astrviktor/golang/src/skillbox_diploma/cmd/data/"
const addr = "127.0.0.1:9999"
const SMSDir = DataDir + "sms.data"
const MMSAddr = "http://" + addr + "/mms"
const VoiceCallDir = DataDir + "voice.data"
const EmailDir = DataDir + "email.data"
const BillingDir = DataDir + "billing.data"
const SupportAddr = "http://" + addr + "/support"
const IncidentAddr = "http://" + addr + "/incident"

func simulatorListenAndServeHTTP() {
	router := mux.NewRouter()

	router.HandleFunc("/mms", handleMMS).Methods("GET")
	router.HandleFunc("/support", handleSupport).Methods("GET")
	router.HandleFunc("/incident", handleIncident).Methods("GET")

	http.ListenAndServe(addr, router)
}

func handleMMS(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(DataDir + "mms.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10)
	if random%5 == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleSupport(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(DataDir + "support.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10)
	if random%5 == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleIncident(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(DataDir + "incident.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10)
	if random%5 == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func StartSimulatorServer() {
	simulatorListenAndServeHTTP()
}
