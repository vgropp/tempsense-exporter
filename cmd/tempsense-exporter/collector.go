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
	collector.readDevices(ch, devices)
}

func (collector *TempsenseCollector) readDevices(ch chan<- prometheus.Metric, devices *hid.HidDevices) {
	for _, device := range devices.Devices {
		collector.readSensors(device, ch)
	}
}

func (collector *TempsenseCollector) readSensors(device hid.Device, ch chan<- prometheus.Metric) {
	numSens := 0
	for {
		numSens++
		data, err := device.ReadSensor()
		if err != nil {
			fmt.Printf("error reading device %v: %v\n", device.GetNum(), err)
			break
		}
		collector.addToMetric(ch, data, device.GetNum())
		if numSens >= int(data.SensorCount) {
			break
		}
	}
}

func (collector *TempsenseCollector) addToMetric(ch chan<- prometheus.Metric, data *hid.Data, deviceNum int) {
	select {
	case ch <- prometheus.MustNewConstMetric(collector.tempsenseMetric, prometheus.GaugeValue, data.Temperature(),
		strconv.Itoa(deviceNum), strconv.Itoa(int(data.SensorCurrent)), data.SensorsIdHex()):
	default:
	}
}
