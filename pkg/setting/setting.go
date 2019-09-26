package setting

import (
	"github.com/go-ini/ini"
)

var (
	Config *ini.File
	RunMode string
	ServerPort string
)

func init()  {
	Config, err = ini.Load("/config/app.ini")
	if err != nil {
		
	}
}