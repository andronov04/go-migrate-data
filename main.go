package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)



// sync loop
// data -> plugin -> mapping -> defaults -> in-memory db
// diff or not diff -> CRUD -> finish

// read from config
// unique key
// callbacks
// cycle get data in values parallel

// DRIVER
// connect to ldap
// - filter
// - and etc
// only get data

// PLUGINS
// data transform
// - get attributes
// - mapping
// - set defaults

// SETTINGS
// - no diff(overwrite)
// - dry-run
// - backup

// - get data from two source parallel
// - iterate data
// -- item transform(mapping, defaults)
// -- save in memory
// --- maybe parallel diff

// SOURCES
// - PostgreSQL
// - AD
// - LDAP
// - MongoDB

// two source
// source ------> target


func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:Fh3fu3IIu@localhost:27017"))
	coll := client.Database("getmobit").Collection("task")


	filter := bson.M{"status": "INITIAL"}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	result, err := coll.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
	fmt.Println(result)

}
