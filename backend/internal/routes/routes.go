package routes

import (
	"encoding/json"
	"fmt"
	"github/Omar-Radwan/backend/internal/mongo"
	"github/Omar-Radwan/backend/internal/redis"
	"github/Omar-Radwan/backend/internal/root"
	"net/http"

	"github.com/gorilla/mux"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", Root).Methods("GET")
	router.HandleFunc("/redis", Redis).Methods("GET")
	router.HandleFunc("/mongo", Mongo).Methods("GET")

	return router
}

type Response struct {
	Value int64
}

func Root(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(Response{root.InsertAndCount(r.URL.Query().Has("increment"))})
	w.Write(b)

}

func Redis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("GET Request Made to /redis")

	b, _ := json.Marshal(Response{redis.InsertAndCount(r.URL.Query().Has("increment"))})
	w.Write(b)

}

func Mongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	b, _ := json.Marshal(Response{mongo.InsertAndCount(r.URL.Query().Has("increment"))})
	w.Write(b)
}
