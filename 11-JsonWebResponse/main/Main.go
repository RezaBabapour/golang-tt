package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Server struct {
	ServerAddr string `json:"serverAddr"`
}

func main() {
	jsonFile, err := os.Open("servers.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	strs := result["servers"].([]interface{})
	servers := make([]interface{}, 0)
	for _, i := range strs {
		server := i.(map[string]interface{})
		servers = append(servers, server["serverAddr"])
	}
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		for _, server := range servers {
			log.Println(r)
			fmt.Fprintln(w, server.(string))
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
