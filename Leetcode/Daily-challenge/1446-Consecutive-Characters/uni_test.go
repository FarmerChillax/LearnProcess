package consecutivecharacters

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestMaxPower(t *testing.T) {
	input := "tourist"
	expected := 1
	result := maxPower(input)
	fmt.Println("result:", result)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("Not Match,", result)
	}
}

func TestMaxPower_2(t *testing.T) {
	input := "leetcode"
	expected := 2
	result := maxPower(input)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("Not Match,", result)
	}
}

func TestMaxPower_3(t *testing.T) {
	input := "abbcccddddeeeeedcba"
	expected := 5
	result := maxPower(input)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("Not Match,", result)
	}
}

func TestMaxPower_4(t *testing.T) {
	input := "triplepillooooow"
	expected := 5
	result := maxPower(input)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("Not Match,", result)
	}
}

func TestMaxPower_5(t *testing.T) {
	input := "hooraaaaaaaaaaay"
	expected := 11
	result := maxPower(input)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("Not Match,", result)
	}
}

func TestMaxPower_6(t *testing.T) {
	input := "aa"
	expected := 2
	result := maxPower(input)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("Not Match,", result)
	}
}
