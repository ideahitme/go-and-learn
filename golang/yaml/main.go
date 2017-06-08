package main

import (
	"fmt"

	"encoding/json"

	"gopkg.in/yaml.v2"
)

type YAMLPerson struct {
	Name string `yaml:"yaml-name"`
}
type JSONPerson struct {
	Name string `json:"json-name"`
}

func main() {
	ObjectToString()
	StringToObject()
	StringToMap()
}

func StringToObject() {
	var yp YAMLPerson
	var err error

	yamlData := `name: yerken`
	err = yaml.Unmarshal([]byte(yamlData), &yp)
	if err != nil {
		panic(err)
	}
	fmt.Println(yp) //empty
	yp = YAMLPerson{}

	yamlData = `yaml-name: yerken`
	err = yaml.Unmarshal([]byte(yamlData), &yp)
	if err != nil {
		panic(err)
	}
	fmt.Println(yp) //{yerken}
	yp = YAMLPerson{}

	yamlData = `{"yaml-name": "yerken"}
`
	err = yaml.Unmarshal([]byte(yamlData), &yp)
	if err != nil {
		panic(err)
	}
	fmt.Println(yp) //{yerken}
	yp = YAMLPerson{}

	jsonData := `{"json-name":"yerken"}`
	jp := JSONPerson{}
	err = yaml.Unmarshal([]byte(jsonData), &jp)
	if err != nil {
		panic(err)
	}
	fmt.Println(jp) //{}

	jsonData = `{"json-name":"yerken"}`
	jp = JSONPerson{}
	err = json.Unmarshal([]byte(jsonData), &jp)
	if err != nil {
		panic(err)
	}
	fmt.Println(jp) //{yerken}
}

func StringToMap() {
	var m map[string]interface{}
	jsonData := `{"name":"yerken"}`
	yaml.Unmarshal([]byte(jsonData), &m)
	fmt.Println(m) //map[name:"yerken"]
}

func ObjectToString() {
	p := &YAMLPerson{"yerken"}
	b, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b)) // yaml-name: yerken

	jp := &JSONPerson{"yerken"}
	b, err = yaml.Marshal(jp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b)) //ignores json tags: name: yerken
}
