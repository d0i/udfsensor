package plugin

import (
	"github.com/d0i/udfsensor"
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"fmt"
)

func init() {
	udf.MustRegisterGlobalUDSFCreator("us_sensor",
		udf.MustConvertToUDSFCreator(udfsensor.CreateSensor))
	fmt.printf("registered us_sensor")
}
	
