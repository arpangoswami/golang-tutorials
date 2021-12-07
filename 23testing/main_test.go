package main

import "testing"

var test = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	err      bool
}{
	{"valid_data", 100.0, 10.0, 10.0, false},
	{"error_data", 100.0, 0.0, 0.0, true},
	{"expected_five", 100.0, 20.0, 5.0, false},
	{"expected_fraction", 1.0, 2.0, 0.5, false},
}

func TestDivison(t *testing.T) {
	for _, tt := range test {
		got, err := Divide(tt.dividend, tt.divisor)
		if tt.err {
			if err == nil {
				t.Error("Expected an error but didn't get one", err.Error())
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("Ecpected %f but got %f", tt.expected, got)
		}
	}
}

// func TestDivide(t *testing.T) {
// 	_, err := Divide(10.0, 1.0)
// 	if err != nil {
// 		t.Error("Got an error when we shouldn't have got")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := Divide(10.0, 0.0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have got")
// 	}
// }
