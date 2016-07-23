package plugin

import (
	"github.com/d0i/udfsensor"
	_ "gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"fmt"
)

func init() {
	udf.MustRegisterGlobalUDSFCreator("us_sensor",
		udf.MustConvertToUDSFCreator(udfsensor.CreateSensor))
	fmt.Println("registered us_sensor")
}
	
