package typeconv2

import "testing"

// TestDollarToInt function tests ...
func TestDollarToInt(t *testing.T) {

	t.Logf("running...")

	test1 := "*"
	test2 := "$9.00"
	test3 := "$9"
	test4 := "$9.01"
	test5 := "$10,900"
	test6 := "$10,900.00"
	test7 := "$10,900.10"
	test8 := "#"
	test9 := "~~"

	result1 := DollarToInt(test1)
	result2 := DollarToInt(test2)
	result3 := DollarToInt(test3)
	result4 := DollarToInt(test4)
	result5 := DollarToInt(test5)
	result6 := DollarToInt(test6)
	result7 := DollarToInt(test7)
	result8 := DollarToInt(test8)
	result9 := DollarToInt(test9)

	t.Logf("test1[%v], result1[%v]", test1, result1)
	t.Logf("test2[%v], result2[%v]", test2, result2)
	t.Logf("test3[%v], result3[%v]", test3, result3)
	t.Logf("test4[%v], result4[%v]", test4, result4)
	t.Logf("test5[%v], result5[%v]", test5, result5)
	t.Logf("test6[%v], result6[%v]", test6, result6)
	t.Logf("test7[%v], result7[%v]", test7, result7)
	t.Logf("test8[%v], result8[%v]", test8, result8)
	t.Logf("test9[%v], result9[%v]", test9, result9)

	t.Logf("exiting")
}

// TestToInt function tests ...
func TestToInt(t *testing.T) {

	t.Logf("running...")
	var iInt int
	var iInt64 int64
	var ifaceInt, ifaceInt64, ifaceNil interface{}

	iInt = 109
	iInt64 = 1929292929390030003

	test3 := "89,983, 899"
	test4 := "#250*"

	ifaceInt = iInt
	ifaceInt64 = iInt64
	ifaceNil = nil

	result1 := ToInt(ifaceInt)
	result2 := ToInt(ifaceInt64)
	result3 := ToInt(test3)
	result4 := ToInt(test4)
	result5 := ToInt(ifaceNil)

	t.Logf("test1[%v], result1[%v]", ifaceInt, result1)
	t.Logf("test2[%v], result2[%v]", ifaceInt64, result2)
	t.Logf("test3[%v], result3[%v]", test3, result3)
	t.Logf("test4[%v], result4[%v]", test4, result4)
	t.Logf("test5[%v], result5[%v]", ifaceNil, result5)

	t.Logf("exiting")
}

// TestToInt64 function tests ...
func TestToInt64(t *testing.T) {

	t.Logf("running...")
	var iInt int
	var iInt64 int64
	var ifaceInt, ifaceInt64, ifaceNil interface{}

	iInt = 109
	iInt64 = 1929292929390030003

	test3 := "89,983, 899"
	test4 := "#250*"

	ifaceInt = iInt
	ifaceInt64 = iInt64
	ifaceNil = nil

	result1 := ToInt64(ifaceInt)
	result2 := ToInt64(ifaceInt64)
	result3 := ToInt64(test3)
	result4 := ToInt64(test4)
	result5 := ToInt64(ifaceNil)

	t.Logf("test1[%v], result1[%v]", ifaceInt, result1)
	t.Logf("test2[%v], result2[%v]", ifaceInt64, result2)
	t.Logf("test3[%v], result3[%v]", test3, result3)
	t.Logf("test4[%v], result4[%v]", test4, result4)
	t.Logf("test5[%v], result5[%v]", ifaceNil, result5)

	t.Logf("exiting")
}

// TestToFloat function tests ...
func TestToFloat(t *testing.T) {

	t.Logf("running...")

	test1 := "100^00 "
	test2 := "49.1"
	test3 := "2.22"
	test4 := "19."

	result1 := ToFloat(test1)
	result2 := ToFloat(test2)
	result3 := ToFloat(test3)
	result4 := ToFloat(test4)

	t.Logf("test1[%v], result1[%v]", test1, result1)
	t.Logf("test2[%v], result2[%v]", test2, result2)
	t.Logf("test3[%v], result3[%v]", test3, result3)
	t.Logf("test4[%v], result4[%v]", test4, result4)

	t.Logf("exiting")
}

// TestToStringDecimal function tests ...
func TestToDecimal(t *testing.T) {

	t.Logf("running...")

	var s1, s2 string
	var if1, if2 interface{}
	var f1, f2 float64

	s1 = "*"
	s2 = "2.21"
	if1 = "1009456.00"
	if2 = nil
	f1 = 1.10
	f2 = 0.99

	result1 := ToDecimal(s1)
	result2 := ToDecimal(s2)
	result3 := ToDecimal(if1)
	result4 := ToDecimal(if2)
	result5 := ToDecimal(f1)
	result6 := ToDecimal(f2)

	t.Logf("test1[%v], result1[%v]", s1, result1)
	t.Logf("test2[%v], result2[%v]", s2, result2)
	t.Logf("test3[%v], result3[%v]", if1, result3)
	t.Logf("test4[%v], result4[%v]", if2, result4)
	t.Logf("test5[%v], result5[%v]", f1, result5)
	t.Logf("test6[%v], result6[%v]", f2, result6)

	t.Logf("exiting")
}

// TestToDollar function tests ...
func TestToDollar(t *testing.T) {

	t.Logf("running...")

	var i1, i2 int
	var i641, i642 int64
	var s1, s2 string
	var f1, f2 float64

	i1 = 10
	i2 = 9990
	i641 = 9
	i642 = 9990900900
	s1 = "^^"
	s2 = "3"
	f1 = 01
	f2 = 100.1

	result1 := ToDollar(i1)
	result2 := ToDollar(i2)
	result3 := ToDollar(i641)
	result4 := ToDollar(i642)
	result5 := ToDollar(s1)
	result6 := ToDollar(s2)
	result7 := ToDollar(f1)
	result8 := ToDollar(f2)

	t.Logf("test1[%v], result1[%v]", i1, result1)
	t.Logf("test2[%v], result2[%v]", i2, result2)
	t.Logf("test3[%v], result3[%v]", i641, result3)
	t.Logf("test4[%v], result4[%v]", i642, result4)
	t.Logf("test5[%v], result5[%v]", s1, result5)
	t.Logf("test6[%v], result6[%v]", s2, result6)
	t.Logf("test7[%v], result7[%v]", f1, result7)
	t.Logf("test8[%v], result8[%v]", f2, result8)

	t.Logf("exiting")
}

// TestToThousand function tests ...
func TestToThousand(t *testing.T) {

	t.Logf("running...")

	var i1, i2 int
	var i641, i642 int64
	var it1, it2 interface{}

	i1 = 10
	i2 = 9990
	i641 = 400400
	i642 = 900900900
	it1 = "100100"
	it2 = nil

	result1 := ToThousand(i1, false)
	result2 := ToThousand(i2, false)
	result3 := ToThousand(i641, true)
	result4 := ToThousand(i642, false)
	result5 := ToThousand(it1, false)
	result6 := ToThousand(it2, false)

	t.Logf("test1[%v], result1[%v]", i1, result1)
	t.Logf("test2[%v], result2[%v]", i2, result2)
	t.Logf("test3[%v], result3[%v]", i641, result3)
	t.Logf("test4[%v], result4[%v]", i642, result4)
	t.Logf("test5[%v], result5[%v]", it1, result5)
	t.Logf("test6[%v], result6[%v]", it2, result6)

	t.Logf("exiting")
}
