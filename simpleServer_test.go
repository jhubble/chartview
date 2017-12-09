package main

import "testing"
import "encoding/json"


func TestReverseToReturnReversedInputString(t *testing.T) {

        count := 2
        resultJson := generate_data(count,"hour",0,100,"memory")

        var actualResult DataPoints
        err := json.Unmarshal([]byte(resultJson),&actualResult)
        if err != nil {
                t.Fatalf("Invalid JSON produced by generate_data %s",resultJson)
        }


        expectedType := "memory"
	if actualResult.Datatype != expectedType {
		t.Fatalf("Expected %s but got %s", expectedType, actualResult.Datatype)
	}

        if len(actualResult.Data) != count {
		t.Fatalf("Expected %d datapoints, but got %d", len(actualResult.Data), count)
        }
}
