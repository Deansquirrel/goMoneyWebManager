package global

import (
	"context"
	"github.com/Deansquirrel/goMoneyWebManager/object"
)

const (
	//PreVersion = "1.0.7 Build20190921"
	//TestVersion = "0.0.0 Build20190101"
	Version   = "0.0.0 Build20190101"
	Type      = "MoneyWebManager"
	SecretKey = "MoneyWebManager"
)

var Ctx context.Context
var Cancel func()

//程序启动参数
var Args *object.ProgramArgs

//系统参数
var SysConfig *object.SystemConfig
