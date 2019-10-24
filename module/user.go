package module

type user struct {
}

//func (user *user) CheckLogin(info *LoginCheckInfo)(*UserInfo,error){
//	if info.CacheId == "" {
//		return user.checkLoginByCacheId(info.CacheId)
//	} else {
//		return user.checkLogin(info.LoginName,info.Password)
//	}
//}

//TODO
func (user *user) CheckLoginByCacheId(id string) (string, error) {
	if id == "123456789" {
		return "123456", nil
	}
	return "", nil
}

//TODO
func (user *user) CheckLogin(name string, pwd string) (string, error) {
	if name == "yuansong" && pwd == "yuansong" {
		return "123456", nil
	}
	return "", nil
}

func (user *user) GetUserInfo(userId string) (string, error) {
	return "", nil
}

//UserId string `json:"userId"`
//UserName string `json:"userName"`
//UserShowName string `json:"userShowName"`
