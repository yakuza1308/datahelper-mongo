package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	//"gopkg.in/mgo.v2"
	. "local_ozone/helper"
)

type msg struct {
	Id    bson.ObjectId `bson:"_id"`
	Msg   string        `bson:"msg"`
	Count int           `bson:"count"`
}

func main() {
	//Create Data
	doc := msg{Id: bson.NewObjectId(), Msg: "Hello from go"}
	fmt.Println("Data Mentah : ")
	fmt.Println(doc)
	//Insert Data
	Save("test", doc)
	//Put Query to select the data that we want to update
	var query []bson.M
	query = append(query, bson.M{"_id": doc.Id})
	newUpdate := bson.M{"$set": bson.M{"msg": "Hello from go [updated]"}}
	//Update Data
	Update("test", bson.M{"$and": query}, newUpdate)

	fmt.Println("-------------------------------------------------")
	fmt.Println("Updated Data :")
	// Select One Row Data
	data, _ := PopulateOneRow("test", bson.M{"$and": query}, nil)
	fmt.Println(data)
	fmt.Println("-------------------------------------------------")
	fmt.Println("List Data :")
	//Select All Data
	dataList, _ := Populate("test", nil, nil, 0, 0)
	for _, d := range dataList {
		fmt.Println(d)
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("Select Column :")
	//Select Msg Column Only
	d := SelectedColumn("msg") //More than one column ? use this : SelectedColumn("msg","count")
	fmt.Println(d)
	dataList, _ = Populate("test", nil, d, 0, 5)
	for _, d := range dataList {
		fmt.Println(d)
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("List Data As Object :")
	var list []msg
	sort := []string{"-msg", "-count"}
	PopulateAsObject(&list, "test", nil, 0, 5, sort...)
	for _, x := range list {
		fmt.Printf("%v | %v\n", x.Msg, x.Count)
	}
	//Aggregate
	fmt.Println("-------------------------------------------------")
	fmt.Println("Aggregate Result :")
	var pipes []bson.M
	pipes = append(pipes, bson.M{"$match": bson.M{"msg": "Hello from go [updated]"}})
	pipes = append(pipes, bson.M{"$group": bson.M{"_id": "$msg"}})
	AggregateResult, _ := Aggregate("test", pipes)
	for _, a := range AggregateResult {
		fmt.Println(a)
	}

}
