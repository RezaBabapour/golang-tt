package main

import (
	"encoding/json"
	"os"
)

func main() {

	file, _ := os.ReadFile("format.json")
	data := make(map[string]interface{})
	_ = json.Unmarshal(file, &data)
	xl := data["xl"].(map[string]interface{})
	sets := xl["sets"].([]interface{})
	set1, set2, set3 := sets[2], sets[1], sets[0]
	sets[0] = set1
	sets[1] = set2
	sets[2] = set3
	xl["sets"] = sets
	data["xl"] = xl
	newData, _ := json.Marshal(data)
	os.WriteFile("format.json",newData,777)
}
