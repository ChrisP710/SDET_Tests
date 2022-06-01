package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// import (
// 	"io/ioutil"
// 	"encoding/json"
// )

type dbDocRows struct {
	total_rows int   `json:"total_rows"`
	offset     int   `json:"offset"`
	rows       dbDoc `json:"dbDoc`
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
	// data := make([]dbDocRows, 0)
	var data map[string]interface{}

	err = json.Unmarshal([]byte(fileRead), &data)
	if err != nil {
		log.Fatalf("Error reading file!: %s", err.Error())
	}

	outputFile, err := json.MarshalIndent(data, "", "  ")

	err = ioutil.WriteFile("test1.json", outputFile, 0644)
}
