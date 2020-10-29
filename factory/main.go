package main

import (
	"fmt"
	"github.com/d3z41k/patterns/factory/factory"
	"reflect"
)

func main() {
	env1 := "production"
	env2 := "development"

	db1 := factory.DatabaseFactory(env1)
	db2 := factory.DatabaseFactory(env2)

	db1.PutData("test", "this is mongodb")
	fmt.Println(db1.GetData("test"))

	db2.PutData("test", "this is sqlite")
	fmt.Println(db2.GetData("test"))

	fmt.Println(reflect.TypeOf(db1).Name())
	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())
}
