package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.Println("start http server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/index", indexHandler)
	mux.HandleFunc("/healthz", healthHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:32000", mux))

}

//index handler
func indexHandler(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, " "))
	}
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	clientAddr := r.RemoteAddr
	w.Write([]byte("hello, From " + clientAddr))
	log.Printf("client: %s, http response code: %d", clientAddr, http.StatusOK)

}

//home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("home page")
	html := "<h1>Http Server</h1><p>hello http server</p>"
	w.Write([]byte(html))
}

//health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["status"] = "ok"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error happened, %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
