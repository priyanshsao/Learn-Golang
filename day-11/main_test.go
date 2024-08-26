package main

import "testing"

// Define a set of test cases for the division function.
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"Valid-Data",12.0,6.0,2.0,false},
	{"Invalid-Data",3.9,0.0,0.0,true},
}
// TestDivision tests the divide function using a table-driven approach (short method).
func TestDivision(t *testing.T) {

	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not got one")
			}
		}else {
			if err != nil {
				t.Error("did not expected an error but got one", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("Expected %f but got %f",tt.expected,got)
		}  
	}
}

// TestDivide tests the divide function with a valid case (long method).
func TestDivide(t *testing.T) {
	_, err := divide(2.0, 3.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

// TestWrongDivide tests the divide function with an invalid case (long method).
func TestWrongDivide(t *testing.T) {
	_, err := divide(2.0, 0.0)
	if err == nil {
		t.Error("did not got an error when we should  have")
	}
}
