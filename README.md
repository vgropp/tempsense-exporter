# tempsense-exporter

This project is intended to provide a prometheus exporter for the DS18B20 based USB HID  Temp Sensor made by Diamex GmbH
(https://www.led-genial.de/USB-Temperatur-Sensor-Tester-fuer-DS18B20-Rev-C)

It was initially inspired by https://github.com/kybernetyk/tempsense

# build

all: `go build -o . ./cmd/...`

tempsense prometheus exporter:
`go build -o . ./cmd/tempsense-exporter`

simple cli utility:
`go build -o . ./cmd/tempsense-cli`

# install

```
go get github.com/vgropp/tempsense-exporter
cd $GOPATH/src/github.com/vgropp/tempsense-exporter
go install ./cmd/...
```

# metrics

- `tempsense_temperature_c` with label device, sensors and serial (`tempsense_temperature_c{device="1",sensor="1",serial="28b3bee62b200127"} 27.8`) 