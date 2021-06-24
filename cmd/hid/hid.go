package hid

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/karalabe/hid"
)

const bufLen = 16

type Data struct {
	SensorCount   byte
	SensorCurrent byte
	Power         byte
	_             byte
	Temp          uint16
	_             uint16
	SensorId      [8]byte
}

type HidDevices struct {
	Devices []HidDevice
}

type HidDevice struct {
	DeviceInfo  hid.DeviceInfo
	Num         int
	DeviceCount int
}

func (data Data) SensorsIdHex() string {
	return hex.EncodeToString(data.SensorId[:])
}

func (data Data) Temperature() float64 {
	return float64(data.Temp) / 10
}

func LookupDevices() (*HidDevices, error) {
	devices := hid.Enumerate(0x16c0, 0x0480)
	if len(devices) == 0 {
		return nil, fmt.Errorf("no temperature sensor found!")
	}
	var hidDevices []HidDevice
	for i, devInfo := range devices {
		hidDevices = append(hidDevices, HidDevice{DeviceInfo: devInfo, Num: i + 1, DeviceCount: len(devices)})
	}
	return &HidDevices{
		Devices: hidDevices,
	}, nil
}

func (dev HidDevice) ReadSensor() (*Data, error) {
	openDevice, err := dev.DeviceInfo.Open()
	if err != nil {
		return nil, fmt.Errorf("unable to open device %v", err)
	}
	defer func(openDevice *hid.Device) {
		err := openDevice.Close()
		if err != nil {
			panic("error closing device")
		}
	}(openDevice)

	buf := make([]byte, bufLen)
	read, err := openDevice.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("could not read from device: %v", err)
	}
	if read < 0 {
		return nil, fmt.Errorf("read nothing from device")
	}
	if read == bufLen {
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
