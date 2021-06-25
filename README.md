# tempsense-exporter

This project is intended to provide a prometheus exporter for the DS18B20 based USB HID  Temp Sensor made by Diamex GmbH
(https://www.led-genial.de/USB-Temperatur-Sensor-Tester-fuer-DS18B20-Rev-C)

As a test util there is a cli version as well, which just outputs all active sensors.

It was initially inspired by https://github.com/kybernetyk/tempsense

# build
## all
`go build -o . ./cmd/...`

## exporter:
`go build -o . ./cmd/tempsense-exporter`

## simple cli utility:
`go build -o . ./cmd/tempsense-cli`

# install

```
go get github.com/vgropp/tempsense-exporter
cd $GOPATH/src/github.com/vgropp/tempsense-exporter
go install ./cmd/...
```

# running

## exporter
```
$GOPATH/bin/tempsense-exporter -h
Usage of /home/tric/tempsense-exporter:
  -address string
        The address to listen on for HTTP requests. (default ":9181")
```

## cli
```
$GOPATH/bin/tempsense-cli
```

# exported metrics

- `tempsense_temperature_c` with label device, sensors and serial (`tempsense_temperature_c{device="1",sensor="1",serial="28b3bee62b200327"} 27.8`) 

```
# HELP tempsense_temperature_c shows current temperature as reported by the ds18b20
# TYPE tempsense_temperature_c gauge
tempsense_temperature_c{device="1",sensor="1",serial="28b3bee62b200327"} 27.4
```