package main

import "testing"

func TestmaxValue(t *testing.T) {
	// arrange
	M := []int{55, 50, 46, 23, 20, 30}
	expection := 55
	//act
	result := maxValue(M)
	//assert

	if result != expection {
		t.Errorf(" error Expected = %d result is = %d", expection, result)
	}
}
func Testlogic(t *testing.T) {
	x := []int{17, 18, 5, 4, 6, 1}
	expection := []int{18, 6, 6, 6, 1, -1}
	//act
	result := maxValue(M)
	//assert

	if result != expection {
		t.Errorf(" error Expected = %d result is = %d", expection, result)
	}
}
