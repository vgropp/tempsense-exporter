package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/mock"
	"github.com/vgropp/tempsense-exporter/cmd/hid"
	"testing"
)

func TestCallingDevices(t *testing.T) {
	testDevice := new(MockDevice)
	testDevice.On("ReadSensor").Return(&hid.Data{SensorCount: 2, SensorCurrent: 2, Temp: 270, SensorId: [8]byte{1, 2, 3, 4, 5, 6, 7, 8}}, nil).Once()
	testDevice.On("ReadSensor").Return(&hid.Data{SensorCount: 2, SensorCurrent: 1, Temp: 270, SensorId: [8]byte{1, 2, 3, 4, 5, 6, 7, 7}}, nil).Once()
	c := make(chan prometheus.Metric)

	NewTempsenseCollector().readSensors(testDevice, c)

	testDevice.AssertNumberOfCalls(t, "ReadSensor", 2)
}

type MockDevice struct {
	mock.Mock
}

func (dev *MockDevice) ReadSensor() (*hid.Data, error) {
	args := dev.Called()
	return args.Get(0).(*hid.Data), nil
}

func (dev *MockDevice) GetNum() int {
	return 1
}
