package main

import (
	"fmt"
)

type keys struct {
	key1 *bool
	key2 *string
}

func (k *keys) assign_values(v1 bool, v2 string) keys {
	// initialize struct
	value1 := v1
	value2 := v2

	k.key1 = &value1
	k.key2 = &value2

	return keys{
		key1: k.key1,
		key2: k.key2,
	}
}

func main() {
	fmt.Println("This is a placeholder")

	// initialize keys struct
	values := keys{}

	// assign values via functions
	final := values.assign_values(true, "Pointer string")

	// print values
	fmt.Printf("%+v\n", final)
}
