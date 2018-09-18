package mongo_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/sarathsp06/gorest/db/crud"
	"github.com/sarathsp06/gorest/db/crud/mongo"
)

type Person struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Phone string `bson:"phone"`
}

func (p Person) String() string {
	d, _ := json.MarshalIndent(p, " ", "\t")
	return string(d)
}

func getConnection() *mongo.DB {
	db, err := mongo.New("localhost", 27017, "example", time.Second*10)
	if err != nil {
		log.Fatal(err)
	}
	crud.SetDefault(db)
	return db
}

func ExampleDB_Insert() {
	db := getConnection()
	defer db.Session.Close()
	_, err := db.C("person").RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Insert(
		"person",
		&Person{"111111", "Ale", 111, "+55 53 8116 9639"},
		&Person{"111112", "Cla", 11, "+55 53 8402 8510"},
	)
	if err != nil {
		log.Fatal(err)
	}

	var result Person
	err = db.C("person").Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Details: %+v", result.Age)

	// OUTPUT:
	// Details: 111
}

func ExampleDB_Get() {
	db := getConnection()
	defer db.Session.Close()
	_, err := db.C("person").RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Insert("person",
		&Person{"111111", "Ale", 10, "+55 53 8116 9639"},
		&Person{"111112", "Cla", 11, "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}
	var result []Person
	err = db.Get("person", Person{Name: "Ale"}, &result, 1, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Details: %+v", result[0].Age)

	// OUTPUT:
	// Details: 10
}

func ExampleDB_Update() {
	db := getConnection()
	defer db.Session.Close()
	_, err := db.C("person").RemoveAll(bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Insert("person", &Person{"111111", "Ale", 10, "+55 53 8116 9639"}, &Person{"111112", "Cla", 11, "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}
	var result Person
	err = db.Update("person", "111111", Person{Name: "Sarath"}, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Details: %+v", result.Name)

	// OUTPUT:
	// Details: Sarath
}

func ExampleDB_Delete() {
	db := getConnection()
	defer db.Session.Close()
	_, err := db.C("person").RemoveAll(bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Insert(
		"person",
		&Person{"111111", "Ale", 10, "+55 53 8116 9639"},
		&Person{"111112", "Cla", 11, "+55 53 8402 8510"},
	)
	if err != nil {
		log.Fatal(err)
	}
	var result Person
	err = db.Delete("person", "111111", &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Details: %+v\n", result.Phone)
	err = db.Delete("person", "111111", &result)
	if err != nil {
		fmt.Println(err)
	}

	// OUTPUT:
	// Details: +55 53 8116 9639
	// Record not Found
}
