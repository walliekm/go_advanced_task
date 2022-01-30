package main

import (
	"fmt"
	"log"
	"week1/dao"

	"github.com/pkg/errors"
)

func main() {
	sqlDAO, err := dao.NewMySQLDao("localhost", 3306, "test", "passwd")
	if err != nil {
		//数据库无法连接，属于不可恢复的异常，触发panic
		panic(err)
	}

	name := "foo"
	uInfo, err := sqlDAO.GetUserInfoByName(name)
	if err != nil {
		//找不到指定的数据
		//此时根据业务需要，如果不影响后续逻辑，则应忽略此错误，然后继续处理后续逻辑
		//如果会影响到后续逻辑，则提前结束，返回结果，本段代码假设为后者进行处理
		if errors.Cause() == dao.ErrDataNotFound {
			fmt.Printf("user %s is not existed")
			return
		}

		//其它错误，记录日志并显示错误，然后结束
		log.Printf("dao query fail:%+v", err)
		fmt.Printf("Server internal error")
		return
	}

	//没有错误，继续处理后续逻辑
	fmt.Printf("Hello %s, welcome\n", uInfo.Name)
}
