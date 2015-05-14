/*==================================================
 基础模型

 Copyright (c) 2015 翱翔大空 and other contributors
==================================================*/

package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"newWoku/conf"
)

var (
	Db *mgo.Database // 数据库连接池
)

func init() {
	//获取数据库连接
	session, err := mgo.Dial(conf.MONGODB_ADDRESS)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	Db = session.DB("woku")
}

type Base struct {
	Collection *mgo.Collection
}

type BaseModel interface {
	Add(obj interface{}) error
	GetsById(lastId string, limit int, obj interface{}) error
	GetsByPage(page int, limit int, obj interface{}) error
	Get(id string, obj interface{}) error
	Update(id string, update map[string]interface{}) error
	Delete(id string) error
	NewObj() interface{}
	NewSlice() interface{}
}

// 新增资源
func (this *Base) Add(obj interface{}) error {
	return this.Collection.Insert(obj)
}

// 获取资源集
// @param {string} id 上一页最后一个id,没有填空
// @param {Int} limit 显示数量
func (this *Base) GetsById(lastId string, limit int, obj interface{}) error {
	if limit == 0 {
		limit = 10
	}

	if limit < 0 || limit > 100 {
		return errors.New("批量查询数量在1-100之间")
	}

	if !bson.IsObjectIdHex(lastId) {
		return this.Collection.Find(nil).Sort("_id").Limit(limit).All(obj)
	} else {
		return this.Collection.Find(bson.M{"_id": bson.M{"$gt": bson.ObjectIdHex(lastId)}}).Sort("_id").Limit(limit).All(obj)
	}
}

// 获取资源集
// @param {int} page 页码
// @param {Int} limit 显示数量
func (this *Base) GetsByPage(page int, limit int, obj interface{}) error {
	if page == 0 {
		page = 1
	}

	if page < 1 || page > 100 {
		return errors.New("页数在1-100之间")
	}

	if limit == 0 {
		limit = 10
	}

	if limit < 0 || limit > 100 {
		return errors.New("批量查询数量在1-100之间")
	}

	return this.Collection.Find(nil).Sort("_id").Skip((page - 1) * limit).Limit(limit).All(obj)
}

// 获取某个资源
// @param {string} id 资源id
func (this *Base) Get(id string, obj interface{}) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("id" + conf.ERROR_TYPE)
	}

	return this.Collection.FindId(bson.ObjectIdHex(id)).One(obj)
}

// 根据id更新某个资源
func (this *Base) Update(id string, update map[string]interface{}) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("id" + conf.ERROR_TYPE)
	}

	return this.Collection.UpdateId(bson.ObjectIdHex(id), update)
}

// 根据id删除某个资源
func (this *Base) Delete(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("id" + conf.ERROR_TYPE)
	}

	return this.Collection.RemoveId(bson.ObjectIdHex(id))
}
