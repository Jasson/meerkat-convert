package mc_json

import (
	"encoding/json"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

type PostMan struct {
	Info interface{} `json:"info"`
	Item []Item      `json:"item"`
}
type Item struct {
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

type Request struct {
	Method string   `json:"method"`
	Header []Header `json:"header"`
	Body   Body     `json:"body"`
	Url    Url      `json:"url"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Body struct {
	Mode    string  `json:"mode"`
	Raw     string  `json:"raw"`
	Options Options `json:"options"`
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

func ReadJson(fileName string) PostMan {

	log.Println("fileName", fileName)

	var postMan PostMan

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("err", err)
	}
	defer file.Close()

	err1 := json.NewDecoder(file).Decode(&postMan)
	if err1 != nil {
		log.Println("error ", err1)
	}
	log.Printf("person%#vf", postMan)

	return postMan
}
