package middlewares

import (
	"fmt"
	"strconv"
	"strings"
)

const is = " is "
const space = " "
const wrong = "I have no idea what you are talking about"
const credits = "Credits"
const does = "Does"

func CalculateHowMuch(input string, romanMap map[string]string, valueMap map[string]float64, key bool) string {
	result := ""
	tempRoman := ""
	tempQuestion := ""
	if strings.Contains(input, " is ") {
		splitted := strings.Split(input, is)

		splittedValue := strings.Split(splitted[1], space)

		for index, value := range splittedValue {
			_, ok := romanMap[value]
			if key == true {
				if ok {
					tempRoman += romanMap[value]
					tempQuestion += value
					if index == len(splittedValue)-2 {
						numeral, err := NumeralConvert(tempRoman)
						if err != nil {
							fmt.Println("Requested number is in invalid format")
						}
						tempQuestion += fmt.Sprintf(" is %d", numeral)
					} else {
						tempQuestion += " "
					}
				}
			} else {
				if ok {
					tempRoman += romanMap[value]
					tempQuestion += value
					tempQuestion += " "
				} else {
					_, ok := valueMap[value]
					if ok {
						if index == len(splittedValue)-2 {
							numeral, err := NumeralConvert(tempRoman)
							if err != nil {
								fmt.Println("Requested number is in invalid format")
								continue
							}
							finalNumeral := float64(numeral) * valueMap[value]

							tempQuestion += value
							tempQuestion += fmt.Sprintf(" is %f %s", finalNumeral, credits)
						} else {
							tempQuestion += " "
						}
					}
				}
			}

		}
		result = tempQuestion
	} else {
		result = wrong
	}
	return result
}

func Compare(input string, romanMap map[string]string, valueMap map[string]float64, key bool) string {
	output := ""
	input1 := ""
	input2 := ""
	tempRoman := ""
	var value1, value2 int

	if strings.Contains(input, " Does ") {
		splitted := strings.Split(input, does)

		if len(splitted) == 2 {
			parts := strings.Split(splitted[1], "has more Credits than ")
			if len(parts) == 2 {
				input1 = parts[0]
				input2 = parts[1]
			}
		}
		splittedValue1 := strings.Split(input1, space)

		for index, value := range splittedValue1 {
			_, ok := romanMap[value]
			if ok {
				tempRoman += romanMap[value]
				if index == len(splittedValue1)-2 {

					value3, err := NumeralConvert(tempRoman)
					if err != nil {
						fmt.Println("Requested number is in invalid format")
					}
					value1 = value3
				}
			}
		}

		splittedValue2 := strings.Split(input2, space)

		for index, value := range splittedValue2 {
			_, ok := romanMap[value]
			if ok {
				tempRoman += romanMap[value]
				if index == len(splittedValue1)-2 {

					value4, err := NumeralConvert(tempRoman)
					if err != nil {
						fmt.Println("Requested number is in invalid format")
					}
					value2 = value4
				}
			}
		}

		if value1 > value2 {
			output += fmt.Sprintf("%s has more credit than %s", input1, input2)
		} else {
			output += fmt.Sprintf("%s has less credit than %s", input1, input2)
		}
	}

	return output
}

func ReturnWrong() string {
	result := wrong
	return result
}

func StoreValueMap(input string, romanMap map[string]string) (string, float64) {
	splitted := strings.Split(input, is)
	splittedRoman := strings.Split(splitted[0], space)

	tempRoman := ""
	key := ""
	value := 0.00

	credits := strings.Split(splitted[1], space)
	creditValue, err := strconv.Atoi(credits[0])
	FailOnError(err, "Failed to convert string to int")

	for _, char := range splittedRoman {
		_, ok := romanMap[char]
		if ok {
			tempRoman += romanMap[char]
		} else {
			key = char
			numeral, err := NumeralConvert(tempRoman)
			if err != nil {
				fmt.Println("error", err)
			}
			value = float64(creditValue) / float64(numeral)
		}

	}
	return key, value
}

func StoreRomanMap(input string) (string, string) {
	var key string = ""
	var value string = ""
	splitted := strings.Split(input, " is ")
	if len(splitted) > 1 {
		roman := strings.TrimSpace(splitted[1])
		key = splitted[0]
		value = roman
	}
	return key, value
}
