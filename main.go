package main

import (
	"fmt"

	as "github.com/aerospike/aerospike-client-go"
)

func main() {
	client, err := as.NewClient("localhost", 3000)
	if err != nil {
		fmt.Println(err)
	}

	key, err := as.NewKey("namespace", "set",
		"key value goes here and can be any supported primitive")

	bin1 := as.NewBin("bin1", "value1")
	bin2 := as.NewBin("bin2", "value2")

	// Write a record
	err = client.PutBins(nil, key, bin1, bin2)

	// Read a record
	record, err := client.Get(nil, key)

	fmt.Println("records:", record)

	client.Close()
}
