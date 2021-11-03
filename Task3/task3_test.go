package main

import "testing"

func TestOpenTextFile(t *testing.T) {
		// arrange
		file := fileText
		//act
		result := openTextFile().bs
		result = string(result)
		//assert

		if result != expection {
			t.Errorf(" error Expected = %d result is = %d", file, result)
}

