# tempsense-exporter

This project is intended to provide a prometheus exporter for the DS18B20 based USB HID  Temp Sensor made by Diamex GmbH
(https://www.led-genial.de/USB-Temperatur-Sensor-Tester-fuer-DS18B20-Rev-C)

It was initially inspired by https://github.com/kybernetyk/tempsense

# build

tempsense prometheus exporter:
`go build -o bin/tempsense-exporter cmd/exporter/tempsense-*`

simple cli utility:
`go build -o bin/tempsense src/cli/tempsense.go`

# metrics

- `tempsense_temperature_c` with label device, sensors and serial (`tempsense_temperature_c{device="1",sensor="1",serial="28b3bee62b200127"} 27.8`) 