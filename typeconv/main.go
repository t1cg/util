package typeconv

import (
	//standard library
	"bytes"
	"errors"
	"strconv"
	"strings"
)

//global
var (
	errorPassedInvalidParameter = errors.New("Passed invalid argument parameter")
)

// ToFloat64 function formats the string and returns a float64 or an error.
func ToFloat64(raw string) (float64, error) {
	//set up replacer; remove space, $, comma, and *
	r := strings.NewReplacer(" ", "", "$", "", ",", "", "*", "")

	//get result
	result := r.Replace(raw)

	var f float64
	var err error

	//convert to float64
	if len(result) == 0 {
		f = 0.00
	} else {
		f, err = strconv.ParseFloat(result, 64)
		if err != nil {
			return 0.00, err
		}
	}

	return f, nil
}

// ToInt function formats the string and returns an int or an error.
func ToInt(raw string) (int, error) {
	//set up replacer; remove space, $, comma, and *
	r := strings.NewReplacer(" ", "", "$", "", ",", "", "*", "")

	//get result
	result := r.Replace(raw)

	var i int
	var err error

	//convert to int
	if len(result) == 0 {
		i = 0
	} else {
		i, err = strconv.Atoi(result)
		if err != nil {
			return 0, err
		}
	}

	return i, nil
}

// DollarToInt function strips all the characters from the dollar amount.
func DollarToInt(raw interface{}) (int, error) {
	var s string
	var err error

	switch t := raw.(type) {
	case int:
		s = strconv.Itoa(t)
	case int64:
		s = strconv.FormatInt(t, 10)
	case float64:
		s = strconv.FormatFloat(t, 'f', 2, 64)
	default:
		s = strings.TrimSpace(raw.(string))
	}

	//set up replacer; remove space, $, comma, *, and period
	var result string
	r := strings.NewReplacer(" ", "", "$", "", ",", "", "*", "", ".", "")

	//TODO: test all conditions
	switch true {
	case len(s) == 0:
		result = "0"
	case strings.Contains(s, "."):
		result = r.Replace(s)
	default:
		result = r.Replace(s) + "00"
	}

	i, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// ToDollar function ...
func ToDollar(number interface{}) (string, error) {
	var s, first, last, final string

	switch t := number.(type) {
	case int:
		s = strconv.Itoa(t)
	case int64:
		s = strconv.FormatInt(t, 10)
	default:
		s = number.(string)
	}

	switch len(s) {
	case 0:
		return "", errorPassedInvalidParameter
	case 1:
		final = "$0.0" + s
	case 2:
		final = "$0." + s
	default:
		first = s[0 : len(s)-2]
		last = s[len(s)-2:]

		temp, err := ToThousand(first, true)
		if err != nil {
			return "", err
		}
		final = temp + "." + last
	}

	return final, nil
}

// ToThousand function formats the number and returns a string with
// thousands comma.
func ToThousand(number interface{}, dollar bool) (string, error) {
	var s string
	var buff bytes.Buffer
	var index int

	switch t := number.(type) {
	case int:
		s = strconv.Itoa(t)
	case int64:
		s = strconv.FormatInt(t, 10)
	case string:
		s = number.(string)
	default:
		return "", errorPassedInvalidParameter
	}

	//add dollar sign if checked
	if dollar {
		buff.WriteByte('$')
	}

	//set the comma index
	index = 3 - ((len(s)) % 3)

	if index == 3 {
		index = 0
	}

	//range through each character
	for _, v := range s {

		if index == 3 {
			buff.WriteRune(',')
			index = 0
		}

		index++
		buff.WriteRune(v)

	}

	return buff.String(), nil

}
