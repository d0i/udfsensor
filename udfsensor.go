package udfsensor

import (
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"math/rand"
	"math"
	"time"
)


// as data source of SensorBee
// XXX ここ，元のsampleにならってdata.Floatにしたけど，生のfloat64とかではNGなの？
type Sensor struct {
	interval time.Duration
	average data.Float
	stddev data.Float // fluctuation
	cycle_depth data.Float // cycle emulated by math.Sin
}

// GenerateStream generates random sensor value
// XXX これの元定義が見つけられない...
func (this *Sensor) GenerateStream(ctx *core.Context, w core.Writer) error {
	// simple sensor
	for {
		ut := float64(time.Now().Unix())/12.0 // no reason for 12.0
		var v data.Float
		v = this.average + this.cycle_depth * data.Float(math.Sin(ut)) + this.stddev * data.Float(rand.NormFloat64())
		t := core.NewTuple(data.Map{
			"type": data.String("temp"),
			"value": v,
		})
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
	stddev := data.Float(0.3)
	cycle_depth := data.Float(5.0)
	if v, ok := params["interval"]; ok {
		i, err := data.ToDuration(v)
		if err != nil {
			return nil, err
		}
		interval = i
	}
	return core.ImplementSourceStop(&Sensor{
		interval: interval,
		average: average,
		stddev: stddev,
		cycle_depth: cycle_depth,
	}), nil
}
