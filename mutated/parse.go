package mutated

import (
	"fmt"
	"regexp"
	"strconv"
)

var matchDate *regexp.Regexp

var nonLeapDays = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

var leapDays = map[int]int{
	1:  31,
	2:  29,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

const (
	dateFormat = `(^\d{2})/(\d{2})/(\d{4}$)`
)

type date struct {
	day   int
	month int
	year  int
}

func init() {
	//DD/MM/YYYY
	matchDate = regexp.MustCompile(dateFormat)
}

func isLeapYear(aCondition, year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func isValidDate(aCondition int, d *date) (bool, error) {

	//The solution needs to cater for all valid dates between
	//01/01/1901 and 31/12/2999.
    switch aCondition {
        // mutatants
        case 18:
            if d.year < 1901 || d.year > 3000 ||
            d.month < 1 || d.month > 12 ||
            d.day < 1 || d.day > 31 {
                return false, fmt.Errorf("the date time must be within the correct time range:01/01/1901 - 31/12/2999")
            }

        case 19:
            if d.year < 1901 && d.year > 2999 ||
            d.month < 1 || d.month > 12 ||
            d.day < 1 || d.day > 31 {
                return false, fmt.Errorf("the date time must be within the correct time range:01/01/1901 - 31/12/2999")
            }

        case 20:
            if d.year < 1901 || false ||
            d.month < 1 || d.month > 12 ||
            d.day < 1 || d.day > 31 {
                return false, fmt.Errorf("the date time must be within the correct time range:01/01/1901 - 31/12/2999")
            }

        // original code
        default:
            if d.year < 1901 || d.year > 2999 ||
            d.month < 1 || d.month > 12 ||
            d.day < 1 || d.day > 31 {
                return false, fmt.Errorf("the date time must be within the correct time range:01/01/1901 - 31/12/2999")
            }
    }

	if isLeapYear(aCondition, d.year) {
        switch aCondition {
            // mutations
            case 21:
                if d.day < leapDays[d.month] {
                    return false, 
                    fmt.Errorf("the day of the month is out of range,it should be between 1 - %d",
                    leapDays[d.month])
                }

            case 22:
                if d.day >= leapDays[d.month] {
                    return false, 
                    fmt.Errorf("the day of the month is out of range,it should be between 1 - %d",
                    leapDays[d.month])
                }

            case 23:
                if d.day > leapDays[d.month+1] {
                    return false, 
                    fmt.Errorf("the day of the month is out of range,it should be between 1 - %d",
                    leapDays[d.month])
                }

            case 24:
                if d.day > leapDays[1] {
                    return false, 
                    fmt.Errorf("the day of the month is out of range,it should be between 1 - %d",
                    leapDays[d.month])
                }

            // original code
            default:
                if d.day > leapDays[d.month] {
                    return false, 
                    fmt.Errorf("the day of the month is out of range,it should be between 1 - %d",
                    leapDays[d.month])
                }
        }
	} else {
		if d.day > nonLeapDays[d.month] {
			return false, fmt.Errorf("the day of the month is out of range,it should be betwen 1- %d", nonLeapDays[d.month])
		}
	}

	return true, nil
}

func parse(aCondition int, str string) (*date, error) {

	var d date

	res := matchDate.FindAllStringSubmatch(str, -1)

	if len(res) == 0 || len(res[0]) != 4 {
		return nil, fmt.Errorf("failed to parse the input time:%s, %s", str, " It has to be DD/MM/YYYY format exactly")
	}

	// matchDate make sure we get the digits string, we don't need to check err here
    switch aCondition {
        case 25: d.day, _ = strconv.Atoi(res[0][0])
        case 26: d.day, _ = strconv.Atoi(res[1][1])
        default: d.day, _ = strconv.Atoi(res[0][1])
    }

    switch aCondition {
        case 27: d.month, _ = strconv.Atoi(res[0][0])
        case 28: d.month, _ = strconv.Atoi(res[2][2])
        default: d.month, _ = strconv.Atoi(res[0][2])
    }

    switch aCondition {
        case 29: d.year, _ = strconv.Atoi(res[0][0])
        case 30: d.year, _ = strconv.Atoi(res[3][3])
        default: d.year, _ = strconv.Atoi(res[0][3])
    }

	_, err := isValidDate(aCondition, &d)

	if err != nil {
		return nil, fmt.Errorf("the input date is invalid:%s", err.Error())
	}

	return &d, nil
}

