package main

import (
	"fmt"
	"publisher_service/pkg/finance"
)

func main() {
	r := finance.StoreDummyData(3)
	fmt.Println(r)
}
