package main

import (
	"fmt"
	"github.com/vgropp/tempsense-exporter/cmd/hid"
	"time"
)

func main() {
	devices, err := hid.LookupDevices()
	if err != nil {
		panic(err)
		return
	}
	for i, device := range devices.Devices {
		fmt.Printf("device found %v: %s[%s]\n", i+1, device.DeviceInfo.Manufacturer, device.DeviceInfo.Product)
	}
	for {
		for _, device := range devices.Devices {
			readDevice(device)
		}
		// the hid device usually blocks anyway, but better wait here just in case to not loop to fast
		time.Sleep(50 * time.Millisecond)
	}
}

func readDevice(device hid.HidDevice) {
	for {
		curSensor, maxSensors := readNextSensor(device)
		if curSensor >= maxSensors {
			break
		}
	}
}

func readNextSensor(device hid.HidDevice) (byte, byte) {
	data, err := device.ReadSensor()
	if err != nil {
		fmt.Printf("error reading device %v: %v\n", device.Num, err)
		return 0, 0
	}
	fmt.Printf("%v: device %v/%v, sensor %v/%v (%v): %.1f°\n",
		time.Now().Format(time.Stamp), device.Num, device.DeviceCount, data.SensorCurrent, data.SensorCount, data.SensorsIdHex(), data.Temperature())
	return data.SensorCurrent, data.SensorCount
}
