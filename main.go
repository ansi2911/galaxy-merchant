package main

import (
	"fmt"
	"strings"

	"prospace-test/middlewares"
)

const many = "how many"
const much = "how much"
const credits = "Credits"
const questionMark = "?"
const does = "Does"
const do = "Is"

func main() {
	valueMap := make(map[string]float64)
	romanMap := make(map[string]string)

	content := middlewares.Read("input.txt")
	splittedInput := strings.Split(content, "\n")

	for _, input := range splittedInput {
		input = strings.TrimSpace(input)

		if strings.Contains(input, questionMark) {
			if strings.Contains(input, much) {
				key := true
				result := middlewares.CalculateHowMuch(input, romanMap, valueMap, key)
				fmt.Println(result)
			} else if strings.Contains(input, many) {
				key := false
				result := middlewares.CalculateHowMuch(input, romanMap, valueMap, key)
				fmt.Println(result)
			} else if strings.Contains(input, does) {
				key := true
				result := middlewares.Compare(input, romanMap, valueMap, key)
				fmt.Println(result)
			} else if strings.Contains(input, do) {
				key := true
				result := middlewares.Compare(input, romanMap, valueMap, key)
				fmt.Println(result)
			} else {
				result := middlewares.ReturnWrong()
				fmt.Println(result)
			}
		} else {
			if strings.Contains(input, credits) {
				key, value := middlewares.StoreValueMap(input, romanMap)
				valueMap[key] = value
			} else {
				key, value := middlewares.StoreRomanMap(input)
				if key == "" && value == "" {
					result := middlewares.ReturnWrong()
					fmt.Println(result)
				} else {
					romanMap[key] = value
				}
			}
		}
	}
}
