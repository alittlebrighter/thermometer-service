package main

import (
	"github.com/alittlebrighter/embd"
	rpi "github.com/alittlebrighter/embd/host/rpi"
	"github.com/alittlebrighter/embd/sensor/mcp9808"
)

func init() {
	rpi.HostInit()
}

type Thermometer struct {
	sensor *mcp9808.MCP9808
}

func NewThermometer() (*Thermometer, error) {
	t := new(Thermometer)

	var err error
	bus := embd.NewI2CBus(1)
	t.sensor, err = mcp9808.New(bus)
	if err != nil {
		return nil, err
	}

	t.sensor.SetShutdownMode(false)
	t.sensor.SetTempResolution(mcp9808.SixteenthC)
	t.sensor.SetTempHysteresis(mcp9808.Zero)

	return t, nil
}

func (t *Thermometer) ReadTemperature() (*mcp9808.Temperature, error) {
	return t.sensor.AmbientTemp()
}

func (t *Thermometer) Shutdown() {
	t.sensor.SetShutdownMode(true)
	embd.CloseI2C()
}
