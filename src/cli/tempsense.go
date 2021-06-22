package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/karalabe/hid"
	"time"
)

const BufLen = 16

type Data struct {
	SensorCount    byte
	SensorsCurrent byte
	Power          byte
	_              byte
	Temp           uint16
	_              uint16
	SensorId       [8]byte
}

func main() {
	devices := hid.Enumerate(0x16c0, 0x0480)
	if len(devices) == 0 {
		panic("no temperature sensor found!")
	}
	for i, device := range devices {
		fmt.Printf("device found %v: %s[%s]\n", i+1, device.Manufacturer, device.Product)
	}
	for {
		for i, device := range devices {
			readDevice(i, device, devices)
		}
		time.Sleep(500 * time.Millisecond)
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
	for {
		curSensor, maxSensors := readNextSensor(open, devNum, len(devices))
		if curSensor >= maxSensors {
			break
		}
	}
}

func readNextSensor(device *hid.Device, devNum int, devCount int) (byte, byte) {
	data, err := readSensor(device)
	if err != nil {
		fmt.Printf("error reading device %v: %v\n", devNum, err)
		return 0, 0
	}
	fmt.Printf("%v: device %v/%v, sensor %v/%v (%v): %.1f°\n",
		time.Now().Format(time.Stamp),
		devNum, devCount,
		data.SensorsCurrent, data.SensorCount,
		hex.EncodeToString(data.SensorId[:]),
		float32(data.Temp)/10)
	return data.SensorsCurrent, data.SensorCount
}

func readSensor(device *hid.Device) (*Data, error) {
	buf := make([]byte, BufLen)
	read, err := device.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("could not read from device: %v", err)
	}
	if read < 0 {
		return nil, fmt.Errorf("read nothing from device")
	}
	if read == BufLen {
		data, err := decode(buf)
		if err != nil {
			return nil, fmt.Errorf("unable to decode %v", err)
		}
		return data, nil
	}
	return nil, fmt.Errorf("read only %v bytes", read)
}

func decode(b []byte) (*Data, error) {
	buf := bytes.NewBuffer(b)
	data := &Data{}
	if err := binary.Read(buf, binary.LittleEndian, data); err != nil {
		return nil, err
	}
	return data, nil
}
