package dayoftheyear

import (
	"log"
	"testing"
)

func TestDayOfYear(t *testing.T) {
	date := "2019-01-09"
	excepted := 9
	result := dayOfYear(date)
	if result != excepted {
		log.Fatalf("not match, result: %v; excepted: %v\n", result, excepted)
	}

}
func TestDayOfYear_2(t *testing.T) {
	date := "2019-02-10"
	excepted := 41
	result := dayOfYear(date)
	if result != excepted {
		log.Fatalf("not match, result: %v; excepted: %v\n", result, excepted)
	}

}
func TestDayOfYear_3(t *testing.T) {
	date := "2003-03-01"
	excepted := 60
	result := dayOfYear(date)
	if result != excepted {
		log.Fatalf("not match, result: %v; excepted: %v\n", result, excepted)
	}

}
func TestDayOfYear_4(t *testing.T) {
	date := "2004-03-01"
	excepted := 61
	result := dayOfYear(date)
	if result != excepted {
		log.Fatalf("not match, result: %v; excepted: %v\n", result, excepted)
	}

}
