package main

import (
    "log"
    "github.com/morus12/dht22"
)

func main() {
    sensor := dht22.New("GPIO_17")
  
    t, err := sensor.Temperature()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Temp: ", t)
  
    h, err := sensor.Humidity()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Humidity:", h)
}
