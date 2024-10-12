package mutated

import (
	"time"
)

func caculate(aConditions int, start *date, end *date) int {

	// UTC timestamp in seconds
	t2 := time.Date(end.year, time.Month(end.month),
        end.day, 0, 0, 0, 0, time.UTC).Unix()

    var t1 int64
    if aConditions == 1 { // mutation 1
	    t1 = time.Date(start.year, time.Month(start.month),
            start.day, 1, 0, 0, 0, time.UTC).Unix()
    } else {
	    t1 = time.Date(start.year, time.Month(start.month),
            start.day, 0, 0, 0, 0, time.UTC).Unix()
    }

	// end date is ealier, swap them
    switch aConditions {
        // mutants
        case 12: if t1 < t2 { t1, t2 = t2, t1 }
        case 13: if !(t1 > t2) { t1, t2 = t2, t1 }
        case 14: if t1 > t2 { t1, t2 = t2, t2 }
        case 15: if t1 > t2 { t1, t2 = t1, t2 }

        // original code
	    default: if t1 > t2 { t1, t2 = t2, t1 }
    }

    var elapsed int
    switch aConditions {
        // mutants
        case 2 : elapsed = int( ( t2 - t1 ) / ( 3600 * 23 ) - 1 )
        case 3 : elapsed = int( ( t2 - t1 ) / ( 3600 * 24 ) - 2 )
        case 4 : elapsed = int( ( t2 - t1 ) / ( 3600 * 24 ) + 1 )
        case 5 : elapsed = int( ( t2 + t1 ) / ( 3600 * 24 ) - 1 )
        case 6 : elapsed = int( ( t2 * t1 ) / ( 3600 * 24 ) - 1 )
        case 7 : elapsed = int( ( t2 / t1 ) / ( 3600 * 24 ) - 1 )

        // original code
        default: elapsed = int( ( t2 - t1 ) / ( 3600 * 24 ) - 1 )
    }

    switch aConditions {
        // mutants
        case 8 : return 1 // return 1 before the function's final return 0
        case 9 : return 0 // return nothing; skip to final return 0
        case 16: if elapsed != 0 { return elapsed }
        case 17: if elapsed < 0 { return elapsed }

        // original behaviour
        default: if elapsed > 0 { return elapsed }
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

    switch aConditions {
        // mutants
        case 10: return caculate(aConditions, endD, startD), nil
        case 11: return caculate(aConditions, startD, startD), nil

        // original behaviour
        default: return caculate(aConditions, startD, endD), nil
    }
}

