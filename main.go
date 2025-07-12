package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/cacher"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Key struct {
	Key string `json:"key"`
}

var ky string

func init() {
	var a Key
	d, _ := os.ReadFile("env.json")
	json.Unmarshal(d, &a)
	os.Setenv("KEY", a.Key)
	ky = os.Getenv("KEY")
}

func getweather(w http.ResponseWriter, req *http.Request) {
	a := mux.Vars(req)
	s := a["city"]
	var tmp map[string]interface{}
	d, _ := cacher.Readcache(s)
	if d == nil {
		arg, _ := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + s + "?unitGroup=us&key=" + ky + "&contentType=json")
		e, _ := io.ReadAll(arg.Body)
		defer arg.Body.Close()
		rd := arg.StatusCode
		if rd != 200 {
			json.NewEncoder(w).Encode("BAD API REQUEST")
			fmt.Println("fck this:-", rd)
			return
		}

		fmt.Println("Start cache")
		json.Unmarshal(e, &tmp)
		cacher.Addcache(tmp, s)
		json.NewEncoder(w).Encode(string(e))
	} else {
		fmt.Println("Got cached")
		json.NewEncoder(w).Encode(d)
	}

}
func main() {

	router := mux.NewRouter()
	router.Use(cacher.Limitmid)
	router.HandleFunc("/add", getweather).Methods("GET")
	log.Fatal(http.ListenAndServe(":5050", router))

}
