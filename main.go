package main

import (
	"fmt"
	"flag"
	"github.com/MichaelS11/go-dht"
)

var (
	pin = flag.String("pin","GPIO4","pin format GPIOXX")
)

func main() {
	flag.Parse()
	err := dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	dht, err := dht.NewDHT(*pin, dht.Fahrenheit, "")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		return
	}

	humidity, temperature, err := dht.Read()
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Printf("humidity: %v\n", humidity)
	fmt.Printf("temperature: %v\n", temperature)
}
