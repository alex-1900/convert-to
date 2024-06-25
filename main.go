package main

import (
	"fmt"
	"github.com/alex-1900/convert-to/module"
	"github.com/alex-1900/convert-to/structs/profile"
	"gopkg.in/yaml.v3"
	"log"
)

var data = `
a: Easy!
b:
  d: [3, 4]
`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	url := "https://cyanmori.yeahfast.com/link/dTgoNRK5tgaeZtdP?clash=1"
	data, err := module.NewReader(url).Contents()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	clash := profile.Clash{}
	err = yaml.Unmarshal(data, &clash)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- t:\n%v\n\n", clash.Proxies)
}
