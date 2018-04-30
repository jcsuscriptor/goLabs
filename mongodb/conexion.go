package main

import (
	"fmt"
	//"log"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

type artist struct {
	Type string
	Origin string
	Name string
	Website string
	Bio string
}
 
func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db = session.DB("foo")
 
	
	var results []artist


    err = db.C("artist").Find(nil).All(&results)
	
	if err != nil {
        // TODO: Do something about the error
    } else {
        fmt.Println("Results All: ", results) 
    }

}




