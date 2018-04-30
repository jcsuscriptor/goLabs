package main

import (
	"fmt"
	"log"
	"math"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

func  GetTotalPages(pageSize, totalCount int) int {
	totalPages := math.Ceil(float64(totalCount) / float64(pageSize))
	if totalPages == 0.0 {
		return 1
	}
	return int(totalPages)
}
 
func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db = session.DB("foo")
 
	var results []bson.M

	//1. Obtener Total
	totalCount,err :=  db.C("artist").Find(nil).Count()
	fmt.Println("totalCount: ", totalCount) 

	if err != nil {
		log.Fatal(err)
		return
	}

	pageSize := 5
	totalPage := GetTotalPages(pageSize,totalCount)
	
	fmt.Println("totalPage: ", totalPage) 
	for i := 0; i < totalPage; i++ {
	
		skip := totalPage * i
		query := db.C("artist").Find(nil).Sort("_id").Skip(skip).Limit(pageSize)
		err = query.All(&results)
	
		if err != nil {
			log.Fatal(err)
			return
		} else {
			
			fmt.Println("Pagina: ", i)
			for j, item := range results {
				fmt.Println("Fila : ", j )
				for key, value := range item {
					fmt.Println(key, value)
				}
				fmt.Println("==============")
			}
		}
    }

	
	
	
}




