package mc_charles_json

import (
	"encoding/json"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

type Item struct {
	Host string `json:"host"`
	Post string  `json:"post"`
	Port int `json:"port"`
	Path string `json:"path"`
	Request Request `json:"request"`
}

type Request struct {
	Header Header `json:"header"`
	Body   Body     `json:"body"`
	Url    Url      `json:"url"`
}

type Header struct {
	Headers []NameValue `json:"headers"`
}
type NameValue struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type Body struct {
	Text    string  `json:"text"`
	Charset     string  `json:"charset"`
}
type Options struct {
	Raw Raw `json:"raw"`
}
type Raw struct {
	Language string `json:"language"`
}

type Url struct {
	Raw  string   `json:"raw"`
	Host []string `json:"host"`
	Port string   `json:"port"`
	Path []string `json:"path"`
}

func ReadJson(fileName string) interface{} {

	log.Println("fileName", fileName)

	var items []Item

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("err", err)
	}
	defer file.Close()

	err1 := json.NewDecoder(file).Decode(&items)
	if err1 != nil {
		log.Println("error ", err1)
	}
	log.Printf("person%#vf", items)

	return items
}
