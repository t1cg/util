package typeconv2

import (
	"bytes"
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

// DollarToInt64 function strips all the characters from the dollar amount.
func DollarToInt64(raw string) int64 {

	i, err := strconv.ParseInt(replace(raw, "dollar"), 10, 64)
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(err.Error())
		i = 0
	}

	return i
}

// ToFloat function converts string to int.
func ToFloat(raw string) float64 {

	i, err := strconv.ParseFloat(replace(raw, "float"), 64)
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(err.Error())
		i = 0.00
	}

	return i
}

// ToInt function converts string to int.
func ToInt(raw string) int {

	i, err := strconv.Atoi(replace(raw, "int"))
	if err != nil {
		i = 0
	}

	return i
}

// ToInt64 function converts string to int.
func ToInt64(raw string) int64 {

	i, err := strconv.ParseInt(replace(raw, "int"), 10, 64)
	if err != nil {
		i = 0
	}

	return i
}

// ToDollar function ...
func ToDollar(number interface{}) string {
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
		final = "$0.00"
	case 1:
		final = "$0.0" + s
	case 2:
		final = "$0." + s
	default:
		first = s[0 : len(s)-2]
		last = s[len(s)-2:]

		final = ToThousand(first, true) + "." + last
	}

	return final
}

// ToThousand function formats the number and returns a string with
// thousands comma.
func ToThousand(number interface{}, dollar bool) string {
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
		return ""
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
		if !strings.Contains(raw, ".") {
			raw = raw + "00"
		}

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
