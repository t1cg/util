package typeconv2

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	//t1cg util library
	"github.com/t1cg/util/apperror"
)

// DollarToInt function strips all the characters from the dollar amount.
func DollarToInt(raw string) int {

	i, err := strconv.Atoi(replace(raw, "dollar"))
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(err.Error())
		i = 0
	}

	return i
}

// ToFloat function converts string to int.
func ToFloat(raw string) float64 {

	f, err := strconv.ParseFloat(replace(raw, "float"), 64)
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(err.Error())
		f = 0.00
	}

	return f
}

// ToInt function converts string to int.
func ToInt(raw string) int {
	i, err := strconv.Atoi(replace(raw, "int"))
	if err != nil {
		i = 0
	}

	return i
}

// ToInt64 function converts string to int64.
func ToInt64(raw string) int64 {

	i, err := strconv.ParseInt(replace(raw, "int"), 10, 64)
	if err != nil {
		i = 0
	}

	return i
}

// ToDollar function expects either an int, int64, float64 or an interface, and
// returns a string with a dollar sign and 2 decimal places.
// The input is assumed to have the decimal stripped previously. For example, 100 is
// understood to be $1.00
// Input type float64 is NOT expected; the code is there just in case.
func ToDollar(raw interface{}) string {
	var s, first, last, final string

	switch t := raw.(type) {
	case int:
		s = strconv.Itoa(t)
	case int64:
		s = strconv.FormatInt(t, 10)
	case float64:
		s = strconv.FormatFloat(t, 'f', 2, 64)
	default:
		s = replace(raw.(string), "dollar")

		if _, err := strconv.Atoi(s); err != nil {
			s = ""
		}
	}

	switch len(s) {
	case 0:
		final = "$0.00"
	case 1:
		final = "$0.0" + s
	case 2:
		final = "$0." + s
	default:
		if strings.Contains(s, ".") {
			final = s
		} else {
			first = s[0 : len(s)-2]
			last = s[len(s)-2:]

			final = ToThousand(first, true) + "." + last
		}
	}

	return final
}

// ToDecimal function expects either a string or a float64, and returns a
// string with 2 decimal position.
// If input is nil, simply return a blank string.
func ToDecimal(raw interface{}) string {
	var s string

	switch t := raw.(type) {
	case float64:
		s = strconv.FormatFloat(t, 'f', 2, 64)
	case interface{}:
		s = replace(raw.(string), "float")
	default:
		s = ""
	}

	if len(s) > 0 {
		var f, l string

		f = s[0 : len(s)-3]
		l = s[len(s)-2:]

		if len(f) > 3 {
			t := ToThousand(f, false)
			fmt.Println("t:", t)

			s = t + "." + l
		}
	}

	return s
}

// ToThousand function expects either int, int64, or an interface; and formats
// the number and returns a string with thousands comma.
func ToThousand(raw interface{}, dollar bool) string {
	var s string
	var buff bytes.Buffer
	var index int

	switch t := raw.(type) {
	case int:
		s = strconv.Itoa(t)
	case int64:
		s = strconv.FormatInt(t, 10)
	case interface{}:
		s = raw.(string)
	//when nil
	default:
		s = ""
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

	return buff.String()

}

// replace function formats the string and returns an int or an error.
func replace(raw string, replaceType string) string {
	var r string

	switch replaceType {
	case "dollar":
		//if !strings.Contains(raw, ".") {
		//	raw = raw + "00"
		//}

		//set up replacer; remove extras
		s := strings.NewReplacer(" ", "", "$", "", ",", "", "*", "", "#", "", ".", "", "%", "")
		r = s.Replace(raw)

	case "float":
		if !strings.Contains(raw, ".") {
			raw = raw + ".00"
		}

		//set up replacer; remove extras
		s := strings.NewReplacer(" ", "", ",", "", "*", "", "#", "", "%", "")
		r = s.Replace(raw)

	case "int":
		s := strings.NewReplacer(" ", "", ",", "", "*", "", "#", "")
		r = s.Replace(raw)

	}

	//return result
	return r

}
