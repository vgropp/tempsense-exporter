package hid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsing(t *testing.T) {
	var buf = []byte{2, 1, 1, 0, 19, 1, 0, 0, 40, 179, 190, 22, 43, 32, 1, 9}
	data, err := parseBuffer(buf, 16)
	assert.NoErrorf(t, err, "there should be no error")
	assert.Equal(t, byte(2), data.SensorCount, "sensors should be 2")
	assert.Equal(t, byte(1), data.SensorCurrent, "current sensor should be 1")
	assert.Equal(t, uint16(275), data.Temp, "temp should be 275")
	assert.Equal(t, float64(27.5), data.Temperature(), "temp should be 27.5")
	assert.Equal(t, "28b3be162b200109", data.SensorsIdHex(), "sensorid should be 28b3be162b200109")
}

func TestShortRead(t *testing.T) {
	var buf = []byte{2, 1, 1, 0, 19, 1, 0, 0, 40, 179, 190, 22, 43, 32, 1}
	_, err := parseBuffer(buf, 15)
	assert.Error(t, err, "there should be an error")
}

func TestEmptyRead(t *testing.T) {
	var buf = []byte{}
	_, err := parseBuffer(buf, 0)
	assert.Error(t, err, "there should be an error")
}

func TestShortBuffer(t *testing.T) {
	var buf = []byte{2, 1, 1, 0, 19, 1, 0, 0, 40, 179, 190, 22, 43, 32, 1}
	_, err := parseBuffer(buf, 16)
	assert.EqualError(t, err, "unable to decode unexpected EOF", "got wrong error")
}
