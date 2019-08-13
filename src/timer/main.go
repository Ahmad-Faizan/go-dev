package main

import (
	"fmt"
	"time"
)

const (
	defaultDate = "2006-Jan-02"
)

func main() {
	deadline, _ := time.Parse(defaultDate, "2019-Sep-01")
	diff := time.Until(deadline)

	days := int(diff.Hours() / 24)
	hrs := int(diff.Hours()) % 24
	mins := int64(diff.Minutes()) % 60
	secs := int64(diff.Seconds()) % 60

	//fmt.Print(diff, "\n")
	fmt.Printf("%v Days, %v Hours, %v Minutes and %v Seconds remaining", days, hrs, mins, secs)

}
