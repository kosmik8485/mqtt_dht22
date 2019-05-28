package main

import (
	"github.com/d2r2/go-dht"
	logger "github.com/d2r2/go-logger"
	"time"
)

var lg = logger.NewPackageLogger("main",
	//logger.DebugLevel,
	logger.InfoLevel,
)

func main() {
	
	defer logger.FinalizeLogger()	
	logger.ChangePackageLogLevel("dht", logger.InfoLevel)
	
	sensorType := dht.DHT22	
	pin := 4
	
	for {
		temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(sensorType, pin, false, 10)
		if err != nil {
			lg.Fatal(err)
		}
		
		lg.Infof("Sensor = %v: Temperature = %v*C, Humidity = %v%% (retried %d times)",
			sensorType, temperature, humidity, retried)
		time.Sleep( time.Second * 10 )
	}
}
