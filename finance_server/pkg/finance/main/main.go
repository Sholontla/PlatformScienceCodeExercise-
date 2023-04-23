package main

import (
	"finance_server/pkg/finance"
	"fmt"
)

func main() {
	r := finance.FinanceDummyData(1)
	fmt.Println(r)
}
