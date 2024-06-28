package main

import (
	"fmt"
	"github.com/alex-1900/convert-to/module"
)

func main() {
	url := "https://cyanmori.yeahfast.com/link/dTgoNRK5tgaeZtdP?clash=1"
	c := module.NewConverterFromUrl(url, "clash", "clash")
	prof, err := c.Convert()
	if err != nil {
		panic(err)
	}
	fmt.Println(prof.String())
}
