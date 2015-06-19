package helper

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
)

var (
	Host string = "localhost"
	DB   string = "Ozone"
)

func GetDb() (*mgo.Session, error) {
	uri := "mongodb://" + Host + "/"
	if uri == "" {
		log.Fatal("no connection string provided")
	}
	sess, err := mgo.Dial(uri)
	if err != nil {
		log.Fatal(err)
	}
	sess.SetSafe(&mgo.Safe{})
	return sess, err
}
func SelectedColumn(columnName ...string) bson.M {
	result := make(bson.M, len(columnName))
	for _, d := range columnName {
		result[d] = 1
	}
	return result

}
func Populate(collectionName string, query map[string]interface{}, column map[string]interface{}, skip int, limit int) ([]bson.D, error) {
	sess, err := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	var result []bson.D
	err = collection.Find(query).Select(column).Skip(skip).Limit(limit).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

type msg struct {
	Id    bson.ObjectId `bson:"_id"`
	Msg   string        `bson:"msg"`
	Count int           `bson:"count"`
}

func PopulateAsPointer(res interface{}, collectionName string, query map[string]interface{}) {
	result := reflect.MakeSlice(reflect.TypeOf(res), 0, 0).Interface().([]msg)
	sess, err := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	err = collection.Find(query).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(result))

}
func PopulateOneRow(collectionName string, query map[string]interface{}, column map[string]interface{}) (bson.D, error) {
	sess, err := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	var result bson.D
	err = collection.Find(query).Select(column).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}
func Save(collectionName string, docs ...interface{}) {
	sess, _ := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	err := collection.Insert(docs...)
	if err != nil {
		log.Fatal(err)
	}
}
func Update(collectionName string, query map[string]interface{}, update map[string]interface{}) {
	sess, _ := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	err := collection.Update(query, update)
	if err != nil {
		log.Fatal(err)
	}
}
func Delete(collectionName string, query map[string]interface{}) {
	sess, _ := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	err := collection.Remove(query)
	if err != nil {
		log.Fatal(err)
	}
}
