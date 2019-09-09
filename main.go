package main

import (
	"flag"
	"github.com/d2r2/go-dht"

	logger "github.com/d2r2/go-logger"
)

var (
	lg = logger.NewPackageLogger("main",
	logger.DebugLevel,
	// logger.InfoLevel,
	)
	pin = flag.Int("pin", 4, "pin")
)	

func main() {
	flag.Parse()
	defer logger.FinalizeLogger()

	lg.Notify("***************************************************************************************************")
	lg.Notify("*** You can change verbosity of output, to modify logging level of module \"dht\"")
	lg.Notify("*** Uncomment/comment corresponding lines with call to ChangePackageLogLevel(...)")
	lg.Notify("***************************************************************************************************")
	// Uncomment/comment next line to suppress/increase verbosity of output
	// logger.ChangePackageLogLevel("dht", logger.InfoLevel)

	// sensorType := dht.DHT11
	sensorType := dht.DHT22
	// sensorType := dht.DHT12
	// Read DHT11 sensor data from specific pin, retrying 10 times in case of failure.
	pin := 4
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(sensorType, pin, false, 10)
	if err != nil {
		lg.Fatal(err)
	}
	// print temperature and humidity
	lg.Infof("Sensor = %v: Temperature = %v*C, Humidity = %v%% (retried %d times)",
		sensorType, temperature, humidity, retried)
}
