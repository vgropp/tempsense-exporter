package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vgropp/tempsense-exporter/cmd/hid"
	"log"
	"strconv"
)

type TempsenseCollector struct {
	tempsenseMetric *prometheus.Desc
}

func NewTempsenseCollector() *TempsenseCollector {
	return &TempsenseCollector{
		tempsenseMetric: prometheus.NewDesc("tempsense_temperature_c",
			"shows current temperature as reported by the ds18b20",
			[]string{"device", "sensor", "serial"}, nil,
		),
	}
}

func (collector *TempsenseCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.tempsenseMetric
}

func (collector *TempsenseCollector) Collect(ch chan<- prometheus.Metric) {
	devices, err := hid.LookupDevices()
	if err != nil {
		log.Print(err)
		return
	}
	for _, device := range devices.Devices {
		numSens := 0
		for {
			numSens++
			data, err := device.ReadSensor()
			if err != nil {
				fmt.Printf("error reading device %v: %v\n", device.Num, err)
				break
			}
			ch <- prometheus.MustNewConstMetric(collector.tempsenseMetric, prometheus.GaugeValue, data.Temperature(),
				strconv.Itoa(device.Num), strconv.Itoa(int(data.SensorCurrent)), data.SensorsIdHex())
			if numSens >= int(data.SensorCount) {
				break
			}
		}
	}
}
