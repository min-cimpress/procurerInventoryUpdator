package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	itemDnaNbr string
	inventory  string
)

func panicOnErr(err error) {
	panic(err)
}

func init() {
	flag.StringVar(&itemDnaNbr, "itemDnaNbr", "11111", "item DNA number")
	flag.StringVar(&inventory, "inventory", "-1", "inventory count")
}

type config struct {
	ProcurerEndpoint string `json:"IBK-endpoint-getprocurer"`
}

func main() {
	// Get config
	file, _ := os.Open("config.test.json")
	decoder := json.NewDecoder(file)
	config := config{}
	err := decoder.Decode(&config)
	if err != nil {
		panicOnErr(err)
	}

	// strings.Replace("oink oink oink", "k", "ky", 2)

	procurerEndpoint := strings.Replace(config.ProcurerEndpoint, "{Handle}", itemDnaNbr, -1)

	fmt.Println(procurerEndpoint)

}
