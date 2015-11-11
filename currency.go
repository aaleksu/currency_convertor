package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

var api_link = "http://api.fixer.io/latest"

type Rate struct {
	RUB float32
	EUR float32
	USD float32
}

type Currency struct {
	Base string
	Rates Rate
}

func main() {
	resp, err := http.Get(api_link)
	
	if err != nil {
		fmt.Printf("\n%s\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("\n%s\n", err)
	}

	var c Currency
	json_err := json.Unmarshal(body, &c)

	if json_err != nil {
		fmt.Printf("\n%s\n", err)
	}

	fmt.Printf("\n%s:", c.Base)
	fmt.Printf("\n------------------")
	fmt.Printf("\nto RUB -> %7.3f", c.Rates.RUB)
	fmt.Printf("\nto USD -> %7.3f\n", c.Rates.USD)
}
