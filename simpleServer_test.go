package main

import "testing"

func TestReverseToReturnReversedInputString(t *testing.T) {
        actualResult := generate_data("hour",0,100,"memory")

        expectedType := "memory"
	if actualResult.datatype != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedType, actualResult.datatype)
	}
}
