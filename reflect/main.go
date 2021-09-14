package main

import (
	"fmt"
	"reflect"
	"strings"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQueryWithReflect(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQueryWithReflect(e)

	i := 90
	createQueryWithReflect(i)

	fmt.Println("Now create query with interface assertion")
	createQueryWithIfaceAssert(o)
	createQueryWithIfaceAssert(e)
	createQueryWithIfaceAssert(i)
}

// createQueryWithIfaceType creates query based on interface assertion
func createQueryWithIfaceAssert(qry interface{}) {
	switch q := qry.(type) {
	case order:
		fmt.Printf("insert into %v values(%d, %d)\n", "order", q.ordId, q.customerId)
	case employee:
		fmt.Printf("insert into %v values(\"%s\", %d, \"%s\", %d, \"%s\")\n",
			"employee", q.name, q.id, q.address, q.salary, q.country)
	default:
		fmt.Println("unsupported type")
	}
}

// createQueryWithReflect creates query based on reflect
// FieldByName returns the struct field with the given name.
// reflect.ValueOf(qry).FieldByName()
func createQueryWithReflect(qry interface{}) {
	// ValueOf returns a new Value initialized to the concrete value
	if reflect.ValueOf(qry).Kind() == reflect.Struct {
		// TypeOf returns the reflection Type that represents the dynamic type of i.
		// fmt.Println(reflect.TypeOf(qry).Name())
		query := fmt.Sprintf("insert into %s values(", reflect.TypeOf(qry).Name())
		for i := 0; i < reflect.ValueOf(qry).NumField(); i++ {
			// Field returns the i'th field of the struct v.
			// FieldByIndex returns the nested field corresponding to index.
			if reflect.ValueOf(qry).Field(i).Kind() == reflect.String {
				query = fmt.Sprintf("%s\"%s\", ", query, reflect.ValueOf(qry).Field(i).String())
			}
			if reflect.ValueOf(qry).Field(i).Kind() == reflect.Int {
				query = fmt.Sprintf("%s%d, ", query, reflect.ValueOf(qry).Field(i).Int())
			}
		}
		query = strings.TrimRight(query, ", ")
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
	} else {
		fmt.Println("unsupported type")
	}
}
