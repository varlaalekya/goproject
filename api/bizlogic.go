package api

import (
	"context"
	"database/sql"

	"apigolang/dataservice"
	"apigolang/model"
)

type IBizLogic interface {
	CreateStudent(s model.Student) error
}

type BizLogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *BizLogic {
	return &BizLogic{DB: db}
}

func (bl *BizLogic) CreateStudent(s model.Student) error {
	if err := dataservice.InsertStudent(context.Background(), bl.DB, s.Name, s.Age, s.Grade); err != nil {
		return err
	}
	return nil
}