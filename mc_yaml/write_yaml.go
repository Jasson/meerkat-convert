package mc_yaml

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)

}

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func Write(writeContext interface{}, fileName string) {
	d, err := yaml.Marshal(&writeContext)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("--- t dump:\n%s\n\n", string(d))

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file.Write(d)
	defer file.Close()

}
