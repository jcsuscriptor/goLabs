package main

import (
	"fmt"
	//"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

 
func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db = session.DB("foo")
 
	var results []bson.M
	err = db.C("artist").Find(nil).All(&results)
	
	if err != nil {
        // TODO: Do something about the error
    } else {
		fmt.Println("Results All: ", results) 
		for i, item := range results {
			// type of i is int
			// type of s is string
			// s == a[i]
			fmt.Println("Fila : ", i )
			for key, value := range item {
				fmt.Println(key, value)
			}
			fmt.Println("==============")
		}
	}
}




