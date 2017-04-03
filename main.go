package main

import (
	"fmt"

	"time"

	as "github.com/aerospike/aerospike-client-go"
)

func main() {
	client, err := as.NewClient("localhost", 3000)
	if err != nil {
		fmt.Println(err)
	}

	// Initialize policy.
	policy := as.NewWritePolicy(0, 0)
	policy.Timeout = 50 * time.Millisecond // 50 millisecond timeout.

	// Write single value.
	key, _ := as.NewKey("test", "myset", "mykey")

	// Write multiple values.
	bin1 := as.NewBin("name", "John")
	bin2 := as.NewBin("age", 25)

	client.PutBins(nil, key, bin1, bin2)

	// Read a record
	record, err := client.Get(nil, key)

	for index, element := range record.Bins {
		fmt.Println(index + ":", element)
	}

	client.Close()
}
