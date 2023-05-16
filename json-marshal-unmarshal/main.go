package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Person struct {
	Name       string `json:"name"`
	Age        string `json:"custom_age"`
	CreditCard string `json:"-"`
}

func main() {
	p := &Person{Name: "tom", CreditCard: "super secret"}
	pBytes, err := json.Marshal(p)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(pBytes))

	// p2AsRawJson := json.RawMessage(`{"name": "mary", "age": 30}`)
	p2AsRawJson, err := ioutil.ReadFile("mary.json")

	var p2 Person
	json.Unmarshal(p2AsRawJson, &p2)
	log.Printf("%+v", p2)

}
