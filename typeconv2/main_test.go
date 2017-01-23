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

// TestToFloat function tests ...
func TestToFloat(t *testing.T) {

	t.Logf("running...")

	test1 := "*"
	test2 := "9.00"
	test3 := "10"
	test4 := "9.01"
	test5 := "10,900"
	test6 := "10,900.00"
	test7 := "1.00"
	test8 := "#"
	test9 := "~~"

	result1 := ToFloat(test1)
	result2 := ToFloat(test2)
	result3 := ToFloat(test3)
	result4 := ToFloat(test4)
	result5 := ToFloat(test5)
	result6 := ToFloat(test6)
	result7 := ToFloat(test7)
	result8 := ToFloat(test8)
	result9 := ToFloat(test9)

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

	test1 := "100 00 "
	test2 := " *250"
	test3 := "89,983, 899"
	test4 := "#"

	result1 := ToInt(test1)
	result2 := ToInt(test2)
	result3 := ToInt(test3)
	result4 := ToInt(test4)

	t.Logf("test1[%v], result1[%v]", test1, result1)
	t.Logf("test2[%v], result2[%v]", test2, result2)
	t.Logf("test3[%v], result3[%v]", test3, result3)
	t.Logf("test4[%v], result4[%v]", test4, result4)

	t.Logf("exiting")
}
