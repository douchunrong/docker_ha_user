package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Hauser struct {
	ObjectId   string
	HauserInfo interface{}
}

func SaveHaUser(object_id string, m interface{}) error {
	session, err := mgo.Dial(beego.AppConfig.String("url")) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("Hauser").C("hauser") //数据库名称
	_, ha_err := db.Upsert(bson.M{"object_id": object_id}, m)

	if ha_err != nil {
		return err
	}

	return nil
}

func FindHaUser(object_id string) (interface{}, error) {
	session, err := mgo.Dial(beego.AppConfig.String("url")) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("Hauser").C("hauser") //如果该集合已经存在的话，则直接返回

	result := make(map[string]interface{})

	err = db.Find(bson.M{"object_id": object_id}).One(&result)

	if err != nil {
		return &result, err
	}

	return &result, nil
}

func DelHaUser(object_id string) error {
	session, err := mgo.Dial(beego.AppConfig.String("url")) //连接数据库

	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("Hauser").C("hauser") //如果该集合已经存在的话，则直接返回

	err = db.Remove(bson.M{"object_id": object_id})

	if err != nil {
		return err
	}

	return nil
}
