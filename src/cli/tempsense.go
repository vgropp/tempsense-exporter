package main

import (
	"encoding/binary"
	"fmt"
	"github.com/karalabe/hid"
	"time"
)

func main() {
	devices := hid.Enumerate(0x16c0, 0x0480)
	if len(devices) == 0 {
		panic("No temperature sensor found!")
	}
	for i, device := range devices {
		fmt.Printf("device found %v: %s => %s\n", i + 1, device.Manufacturer, device.Product)
	}
	for {
		for i, device := range devices {
			readDevice(i, device, devices)
		}
		time.Sleep(1 * time.Second)
	}
}

func readDevice(i int, device hid.DeviceInfo, devices []hid.DeviceInfo) {
	devNum := i + 1
	open, err := device.Open()
	if err != nil {
		fmt.Printf("unable to open device %v: %v", devNum, err)
		return
	}
	defer func() {
		err = open.Close()
		if err != nil {
			fmt.Printf("unable to close device %v: %v\n", devNum, err)
		}
	}()
	buf := make([]byte, 64)
	read, err := open.Read(buf)
	if err != nil {
		fmt.Printf("could not read from device %v\n", devNum)
		return
	}
	if read < 0 {
		fmt.Printf("read nothing from device %v\n", devNum)
		return
	}
	if read == 64 {
		fmt.Printf("device %v/%v, sensors %v/%v: %.1f\n", devNum, len(devices), buf[1], buf[0], float32(binary.LittleEndian.Uint16(buf[4:]))/10)
	} else {
		fmt.Printf("device %v/%v: read %v bytes", devNum, len(devices), read)
	}
}