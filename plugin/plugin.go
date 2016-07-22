package plugin

import (
	"gohan.to/doi/udfsensor"
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
)

func init() {
	udf.MustRegisterGlobalUDSFCreator("us_sensor",
		udf.MustConvertToUDSFCreator(udfsensor.CreateSensor))
}
	
