package main

import "fmt"

var romanMap = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func main() {
	result := decode("MMVIII")
	fmt.Println(result)
}

func decode(roman string) int {
	sum := 0
	for i := 0; i < len(roman); i++ {
		if val, exists := getValue(roman, i); exists {
			sum += val
			i++
		} else {
			sum += romanMap[string(roman[i])]
		}
	}
	return sum
}

func getValue(roman string, i int) (int, bool) {
	if i+1 < len(roman) && romanMap[string(roman[i])+string(roman[i+1])] > 0 {
		return romanMap[string(roman[i])+string(roman[i+1])], true
	}
	return 0, false
}
