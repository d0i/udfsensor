package udfsensor

import (
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"time"
)


// as data source of SensorBee
type Sensor struct {
	interval time.Duration
	average data.Float
}

// GenerateStream generates random sensor value
// XXX これの元定義が見つけられない...
func (this *Sensor) GenerateStream(ctx *core.Context, w core.Writer) error {
	// first, this is a broken sensor
	for {
		t := core.NewTuple(data.Map{"value": this.average})
		if err := w.Write(ctx, t); err != nil {
			return err
		}
		time.Sleep(this.interval)
	}
}

// stop GenerateStream -- do nothing here (ImplementSourceStop covers it)
func (this *Sensor) Stop(ctx *core.Context) error {
	return nil
}

func CreateSensor(ctx *core.Context, ioParams *bql.IOParams, params data.Map) (core.Source, error) {
	interval := 100 * time.Millisecond
	average := data.Float(28.5)
	if v, ok := params["iterval"]; ok {
		i, err := data.ToDuration(v)
		if err != nil {
			return nil, err
		}
		interval = i
	}
	return core.ImplementSourceStop(&Sensor{
		interval: interval,
		average: average,
	}), nil
}
