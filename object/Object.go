package object

import "time"

type UserInfo struct {
	FId          string
	FName        string
	FShowName    string
	FAddTime     time.Time
	FLastUpdate  time.Time
	FIsForbidden int
}

type LoginInfo struct {
	FId         string
	FLoginName  string
	FAddTime    time.Time
	FLastUpdate time.Time
}

//
//type LoginCheckInfo struct {
//	LoginName string `json:"loginName"`
//	Password string `json:"password"`
//	CacheId string `json:"cacheId"`
//}
//
//type UserInfo struct {
//	UserId string `json:"userId"`
//	UserName string `json:"userName"`
//	UserShowName string `json:"userShowName"`
//	CacheId string `json:"cacheId"`
//}
