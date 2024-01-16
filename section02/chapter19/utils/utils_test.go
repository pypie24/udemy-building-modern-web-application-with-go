package utils

import "testing"


var testCases = []struct{
	name string
	dividend float32
	divisor float32
	expected float32
	isError bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 10.0, 0.0, 0.0, true},
}

func TestDivide(t *testing.T){
	for _, testCase := range testCases{
		got, err := Divide(testCase.dividend, testCase.divisor)
		if testCase.isError && err == nil{
			t.Error("expected the error but did not get one")
		} else if !testCase.isError && err != nil{
			t.Error("unexpected the error but got one")
		}

		if testCase.expected != got{
			t.Errorf("expected the %f, but got the %f", testCase.expected, got)
		}

	}
}