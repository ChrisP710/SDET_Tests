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
	read1()
	// read2()
}

func read1() {
	fileRead, err := ioutil.ReadFile("cpdata.json")
	if err != nil {
		log.Fatalf("Error reading file!: %s", err.Error())
	}
	// Used for debugging
	// fmt.Println("hello", fileRead)

	// Used to hold content after unmarshaling file
	var data map[string]interface{}

	err = json.Unmarshal([]byte(fileRead), &data)
	if err != nil {
		log.Fatalf("Error reading file!: %s", err.Error())
	}

	// Used to test output before creating for loop
	// Grabs key "doc" from nested struct dbDoc
	// key := data["rows"].([]interface{})[20].(map[string]interface{})["doc"]
	// fmt.Println(key)

	// Used to temporarily store value of each array before marshalling and writing it to a new file.
	// Note: 311 Unused aka Null
	keyTemp := make([]interface{}, 688)

	for i := 0; i < len(keyTemp); i++ {
		// Deletes "_rev" fields in document
		delete(data["rows"].([]interface{})[i].(map[string]interface{})["doc"].(map[string]interface{}), "_rev")

		// Test variable I used to test the output and see if "_rev" was being returned in array
		// keyTemp[i] = data["rows"].([]interface{})[i].(map[string]interface{})["doc"].(map[string]interface{})["_rev"]

		// Grabs just the "doc" struct
		keyTemp[i] = data["rows"].([]interface{})[i].(map[string]interface{})["doc"]

		// Marshals file and indents in typical Json Format
		outputFile, err := json.MarshalIndent(keyTemp, "", "  ")
		if err != nil {
			log.Fatalf("Error reading file!: %s", err.Error())
		}
		// Writes to new file
		err = ioutil.WriteFile("doc9.json", outputFile, 0644)
		if err != nil {
			log.Fatalf("Error reading file!: %s", err.Error())
		}
	}

}

// func read2() {
// 	fileRead, err := ioutil.ReadFile("doc1.json")
// 	if err != nil {
// 		log.Fatalf("Error reading file!: %s", err.Error())
// 	}
// 	// fmt.Println("hello", fileRead)

// 	var data []interface{}

// 	err = json.Unmarshal([]byte(fileRead), &data)
// 	if err != nil {
// 		log.Fatalf("Error reading file!: %s", err.Error())
// 	}
// 	// fmt.Println("hello", data[2])

// 	// var data2 map[string]interface{}

// 	// Grabs key "doc" from nested struct dbDoc
// 	// key := data["_rev"]
// 	// defer fmt.Println(key, "hello")

// 	// Note: 311 Unused aka Null
// 	// keyTemp := make([]interface{}, 1000)

// }
