package plugin

import (
	"github.com/d0i/udfsensor.git"
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
)

func init() {
	udf.MustRegisterGlobalUDSFCreator("us_sensor",
		udf.MustConvertToUDSFCreator(udfsensor.CreateSensor))
}
	
