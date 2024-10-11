package mutated

import (
	"time"
)

func caculate(aConditions int, start *date, end *date) int {

	//UTC timestamp in seconds
	t1 := time.Date(start.year, time.Month(start.month), start.day, 0, 0, 0, 0, time.UTC).Unix()
	t2 := time.Date(end.year, time.Month(end.month), end.day, 0, 0, 0, 0, time.UTC).Unix()

	//end date is ealier, swap them
	if t1 > t2 {
		t1, t2 = t2, t1
	}

	elapsed := int((t2-t1)/(3600*24) - 1)

	if elapsed > 0 {
		return elapsed
	}

	return 0
}

// Elapsed caculate the eclapsed days
// start and end are within range: 01/01/1901 and 31/12/2999.
func Elapsed(aConditions int, start, end string) (int, error) {

	startD, err := parse(aConditions, start)

	if err != nil {
		return -1, err
	}

	endD, err := parse(aConditions, end)

	if err != nil {
		return -1, err
	}

	return caculate(aConditions, startD, endD), nil
}

