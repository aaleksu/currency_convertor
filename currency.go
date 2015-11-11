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

	fmt.Printf("\nBase currency: %s", c.Base)
	fmt.Printf("\n------------------\n")
	fmt.Printf("\nRates:")
	fmt.Printf("\nto RUB -> %f", c.Rates.RUB)
	fmt.Printf("\nto USD -> %f\n", c.Rates.USD)
}
