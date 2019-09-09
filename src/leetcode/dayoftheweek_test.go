package leetcode

import (
	"testing"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	thisTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return thisTime.Weekday().String()
}

func TestDayOfTheWeek(t *testing.T) {

	for i := 1; i < 15; i++ {
		weekDay := dayOfTheWeek(i, 9, 2019)
		t.Logf("9.%d ------------> %s\n", i, weekDay)
	}

}
