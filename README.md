# tempsense-exporter

This project is intended to provide a prometheus exporter for the DS18B20 based USB HID  Temp Sensor made by Diamex GmbH
(https://www.led-genial.de/USB-Temperatur-Sensor-Tester-fuer-DS18B20-Rev-C)

As a test util there is a cli version as well, which just outputs all active sensors.

It was initially inspired by https://github.com/kybernetyk/tempsense

# build

## dependencies

you need a recent golang build environment, gcc and linux headers (for karalabe/hid) to build the tempsense-exporter.

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

## docker

Run the exporter in docker might need some tweaking. First you have to expose the usb device by adding `--device=/dev/bus/usb`. Additional the user in docker needs access to the device file. The default is gid 46 (group plugdev in debian). Override the default gid 46 by adding `-u 1000:<gid>` to your docker run command. The exporter is reachable via port 9181, override this with `-p <port>:9181`

Simple example: 
`docker run --device=/dev/bus/usb -p 9181:9181 ghcr.io/vgropp/tempsense-exporter:latest`

Full example:
`docker run --device=/dev/bus/usb -p 9182:9181 -u 1000:46 ghcr.io/vgropp/tempsense-exporter:latest`

# exported metrics

- `tempsense_temperature_c` with label device, sensors and serial (`tempsense_temperature_c{device="1",sensor="1",serial="28b3bee62b200327"} 27.8`) 

```
# HELP tempsense_temperature_c shows current temperature as reported by the ds18b20
# TYPE tempsense_temperature_c gauge
tempsense_temperature_c{device="1",sensor="1",serial="28b3bee62b200327"} 27.4
```
