package common

import (
	"fmt"
	"super_crud/src/config"
)

func GenerateHostURL(short_id string) string{
	host := config.HOST
	port := config.PORT
	version := "/api/v1/file/"
	serverUrl := fmt.Sprintf("%v:%v%v%v",host, port,version, short_id)
	return serverUrl
}