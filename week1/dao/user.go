package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

type UserInfo struct {
	Id     int
	Name   string
	Passwd string
}

func (this *sqlDAO) GetUserInfoByName(name string) (*UserInfo, error) {
	var uInfo UserInfo
	if err := this.db.QueryRow("SELECT id,name,passwd FROM user WHERE name=?", name).Scan(&uInfo.Id, &uInfo.Name, &uInfo.Passwd); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrDataNotFound
		}

		return nil, errors.Wrap(err, "GetUserInfoByName fail")
	}

	return &uInfo, nil
}
