package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type dbDocRows struct {
	total_rows int   `json:"total_rows"`
	offset     int   `json:"offset"`
	rows       dbDoc `json:"rows"`
}

type dbDoc struct {
	id    string                 `json: "id"`
	key   string                 `json: "key"`
	value valueNest              `json: "value"`
	doc   map[string]interface{} `json: "doc"`
}

type valueNest struct {
	rev string `json: "rev"`
}

// type docNest struct {
// 	_id         string `json:"_id"`
// 	_rev        string `json:"_rev"`
// 	source      string `json:"source"`
// 	redirecturl string `json:"redirecturl"`
// 	docType     string `json:"type"`
// 	whitelabel  string `json:"whitelabel"`
// 	created     string `json:"created"`
// 	modified    string `json:"modified"`
// }

func main() {
	fileRead, err := ioutil.ReadFile("cpdata.json")
	if err != nil {
		log.Fatalf("Error reading file!: %s", err.Error())
	}
	// fmt.Println("hello", fileRead)

	var data map[string]interface{}

	err = json.Unmarshal([]byte(fileRead), &data)
	if err != nil {
		log.Fatalf("Error reading file!: %s", err.Error())
	}

	// Grabs key "doc" from nested struct dbDoc
	// key := data["rows"].([]interface{})[20].(map[string]interface{})["doc"]
	// fmt.Println(key)

	// Note: 311 Unused aka Null
	keyTemp := make([]interface{}, 1000)

	for i := 0; i < 1000; i++ {
		keyTemp[i] = data["rows"].([]interface{})[i].(map[string]interface{})["doc"]
		outputFile, err := json.MarshalIndent(keyTemp, "", "  ")
		if err != nil {
			log.Fatalf("Error reading file!: %s", err.Error())
		}

		err = ioutil.WriteFile("testtest2.json", outputFile, 0644)
		if err != nil {
			log.Fatalf("Error reading file!: %s", err.Error())
		}
	}

	// outputFile, err := json.MarshalIndent(key, "", "  ")
	// if err != nil {
	// 	log.Fatalf("Error reading file!: %s", err.Error())
	// }

	// err = ioutil.WriteFile("test999.json", outputFile, 0644)
	// if err != nil {
	// 	log.Fatalf("Error reading file!: %s", err.Error())
	// }
}
