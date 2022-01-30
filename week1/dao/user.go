package dao

type UserInfo struct {
	Id     int
	Name   string
	Passwd string
}

func (this *mysqlDAO) GetUserInfoByName(name string) (uInfo *UserInfo, err error) {
	return
}
