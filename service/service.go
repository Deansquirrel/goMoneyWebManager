package service

import (
	"github.com/Deansquirrel/goMoneyWebManager/global"
	"github.com/Deansquirrel/goMoneyWebManager/webService"
	log "github.com/Deansquirrel/goToolLog"
)

//启动服务内容
func StartService() error {
	log.Debug("Start Service")
	defer log.Debug("Start Service Complete")

	go func() {
		ws := webService.NewWebServer(global.SysConfig.Iris.Port)
		ws.StartWebService()
	}()

	//go func() {
	//	//==========================================================================
	//	goServiceSupportHelper.InitParam(&goServiceSupportHelper.Params{
	//		HttpAddress:   fmt.Sprintf("http://127.0.0.1:%d", global.SysConfig.Iris.Port),
	//		ClientType:    global.Type,
	//		ClientVersion: global.Version,
	//		Ctx:           global.Ctx,
	//		Cancel:        global.Cancel,
	//	})
	//	goServiceSupportHelper.SetOtherInfo(
	//		repository.NewCommon().GetLocalDbConfig(),
	//		0,
	//		goServiceSupportHelper.SVRNONE)
	//	//==========================================================================
	//}()

	//go func() {
	//	for {
	//		err := goToolCron.AddFunc("ClearJobRecord",
	//			global.ClearJobRecordCron,
	//			worker.NewWorker().ClearJobRecord, func(v interface{}) {
	//				log.Error(fmt.Sprintf("panicHandle: %s", v))
	//			})
	//		if err != nil {
	//			log.Error(fmt.Sprintf("add func ClearJobRecord err: %s", err.Error()))
	//			time.Sleep(time.Minute)
	//			continue
	//		}
	//		break
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		err := goToolCron.AddFunc("ClearInvalidHeartBeat",
	//			global.ClearInvalidHeartBeatCron,
	//			worker.NewWorker().ClearInvalidHeartBeat, func(v interface{}) {
	//				log.Error(fmt.Sprintf("panicHandle: %s", v))
	//			})
	//		if err != nil {
	//			log.Error(fmt.Sprintf("add func ClearInvalidHeartBeat err: %s", err.Error()))
	//			time.Sleep(time.Minute)
	//			continue
	//		}
	//		break
	//	}
	//}()

	return nil
}
