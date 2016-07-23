package plugin

import (
	"github.com/d0i/udfsensor"
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	_ "gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"fmt"
)

func init() {
        bql.MustRegisterGlobalSourceCreator("us_sensor",
		bql.SourceCreatorFunc(udfsensor.CreateSensor))
	fmt.Println("registered us_sensor")
}
	
