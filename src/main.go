package main

import (
	"fmt"
	"github.com/HiChen85/hc-go-master/src/basic_syntax"
	"time"
)

func main() {
	var invalidWeekDay time.Weekday = 8
	fmt.Println("invalid: ", invalidWeekDay.String())

	fmt.Println(basic_syntax.SATURDAY)
}
