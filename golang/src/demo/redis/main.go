package main

import (
	"fmt"
	"simbaGolang/demo/redis/rds"
)

func main() {
	a1, err := rds.Set("1", "123")
	fmt.Println(a1, err)

	a2, err := rds.Get("1")
	fmt.Println(a2, err)
}
