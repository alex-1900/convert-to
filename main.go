package main

import (
	"fmt"
	"github.com/alex-1900/convert-to/module"
	"github.com/alex-1900/convert-to/structs/profile"
	"gopkg.in/yaml.v3"
	"log"
)

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
