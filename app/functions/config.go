package functions

import (
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
)

func GetConfig() (*commonConfig.Settings, error) {
	logHandler.InfoLogger.Println("Configuration retrieval is not implemented yet")
	return commonConfig.Get(), nil
}
