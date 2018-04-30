package main

import (
	
	"time"
	"fmt"
	//"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database


type ElmahError struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Application        string        `bson:"application" json:"application"`
	Host  string        `bson:"host" json:"host"`
	Type  string        `bson:"type" json:"type"`
	Message  string        `bson:"message" json:"message"`
	Source  string        `bson:"source" json:"source"`
	Detail  string        `bson:"detail" json:"detail"`
	User  string        `bson:"user" json:"user"`
	Time    time.Time      `bson:"time" json:"time"`
}
 
func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db = session.DB("elmah-go")
 
	
	var results []ElmahError


    err = db.C("Elmah").Find(nil).All(&results)
	
	if err != nil {
        // TODO: Do something about the error
    } else {
        //fmt.Println("Results All: ", results) 
    }

	//Find Datetime - UTE
	fromDate := time.Date(2017, 9, 03, 17, 0, 0, 0, time.UTC)
	toDate :=  time.Date(2017, 9, 03, 18, 0, 0, 0, time.UTC)
 
	var resultsFilter []ElmahError
	err =  db.C("Elmah").Find(
		bson.M{
			"time": bson.M{
				"$gt": fromDate,
				"$lt": toDate,
			},
		}).All(&resultsFilter)
	
	
	if err != nil {
		fmt.Println("Error   : ", err.Error()) 
	} else {
		fmt.Println("Results Filter  (Time UTE): ", len(resultsFilter))
		for _, item := range resultsFilter {
			fmt.Println("Time : ",item.Time)
		}
	}

	//Find Datetime -
	fromDate = time.Date(2017, 9, 03, 12, 0, 0, 0, time.Local)
	toDate =  time.Date(2017, 9, 03, 13, 0, 0, 0, time.Local)

	 
	err =  db.C("Elmah").Find(
		bson.M{
			"time": bson.M{
				"$gt": fromDate,
				"$lt": toDate,
			},
		}).All(&resultsFilter)


	if err != nil {
		fmt.Println("Error   : ", err.Error()) 
	} else {
		fmt.Println("Results Filter (Time Local): ", len(resultsFilter))
		for _, item := range resultsFilter {
			fmt.Println("Time : ",item.Time)
		}
	}
	

}
