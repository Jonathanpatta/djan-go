package djan_go

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type DataModel[T any] struct {
	Db        *gorm.DB
	Data      T
	DataArray []T
}

func RegisterDataModel[T any](data T, c *DataModelConfig) *DataModel[T] {
	if c.GlobalConfig.Debug {
		c.GlobalConfig.GormDb.AutoMigrate(&data)
	}
	return &DataModel[T]{
		Data: data,
		Db:   c.GlobalConfig.GormDb,
	}
}

func (d *DataModel[T]) Get(id string) (T, error) {
	var data T
	err := d.Db.First(&data, "id = ?", id).Error
	return data, err
}

func (d *DataModel[T]) Post(data T) (T, error) {
	st := time.Now()
	err := d.Db.Create(data).Error
	fmt.Println(time.Now().UnixMilli() - st.UnixMilli())
	fmt.Println(d.Db.Error)

	return data, err
}

func (d *DataModel[T]) Put(data T) (T, error) {
	err := d.Db.Save(data).Error
	return data, err
}

func (d *DataModel[T]) Delete(id string) (T, error) {
	data, err := d.Get(id)
	if err != nil {
		return data, err
	}
	err = d.Db.Delete(&d.Data, "id = ?", id).Error
	return data, err
}

func (d *DataModel[T]) List() ([]T, error) {
	var res []T
	err := d.Db.Find(&res).Error
	fmt.Println(res)
	return res, err
}
