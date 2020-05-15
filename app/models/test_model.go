/**
  create by yy on 2020/5/15
*/

package models

import "github.com/jinzhu/gorm"

type TestModel struct {
	Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (t *TestModel) GetDB() *gorm.DB {
	return getTable(PDB, t.TableName())
}

func (t *TestModel) TableName() string {
	return "test"
}

func (t *TestModel) HasTable() bool {

	if PDB == nil {
		return true
	}

	return PDB.HasTable(t.TableName())
}

func (t *TestModel) CreateTable() error {
	db := t.GetDB()
	err := db.CreateTable(t).Error
	if err != nil {
		return err
	}
	return nil
}
