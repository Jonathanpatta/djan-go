package djan_go

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type GormDataModel[T any] struct {
	Db *gorm.DB
}

func NewGormDataModel[T any](c *DataModelConfig) *GormDataModel[T] {
	var data T
	if c.GlobalConfig.Debug {
		c.GlobalConfig.GormDb.AutoMigrate(&data)
	}
	return &GormDataModel[T]{
		Db: c.GlobalConfig.GormDb,
	}
}

func (d *GormDataModel[T]) Get(id string) (T, error) {
	var data T
	err := d.Db.First(&data, "id = ?", id).Error
	return data, err
}

func (d *GormDataModel[T]) Post(data T) (T, error) {
	st := time.Now()
	err := d.Db.Create(data).Error
	fmt.Println(time.Now().UnixMilli() - st.UnixMilli())
	fmt.Println(d.Db.Error)

	return data, err
}

func (d *GormDataModel[T]) Put(data T) (T, error) {
	err := d.Db.Save(data).Error
	return data, err
}

func (d *GormDataModel[T]) Delete(id string) (T, error) {
	data, err := d.Get(id)
	if err != nil {
		return data, err
	}
	err = d.Db.Delete(&data, "id = ?", id).Error
	return data, err
}

func (d *GormDataModel[T]) List() ([]T, error) {
	var res []T
	err := d.Db.Find(&res).Error
	return res, err
}
