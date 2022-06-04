/*
@Date : 2022/6/3 15:00
@Author : cirss
*/
package service

import (
	"github.com/chris1678/go-run/orm"
)

type Service struct {
	Msg   string
	MsgID string
}

func (s Service) GetOne(req interface{}, data interface{}) error {
	var err error
	err = orm.Db().Where(req).First(data).Error
	if err != nil {
		return err
	}
	return nil
}
