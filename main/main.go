package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func envToJson() []byte {
	envmap := make(map[string]string)
	for _, e := range os.Environ() {
		parts := strings.Split(e, "=")
		envmap[parts[0]] = parts[1]
	}
	data, err := json.Marshal(envmap)
	if err != nil {
		panic(err)
	}
	return data
}

func main() {
	log.Println(string(envToJson()))
	handler := HttpHandler{}
	_ = http.ListenAndServe("0.0.0.0:9000", handler)
}

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write(envToJson())
}
