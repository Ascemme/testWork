package main

import "testing"

func TestUsersInput(t *testing.T) {
	// arrange

	expection = UsersInput().allUsers
	//act
	result = UsersInput()

	//assert

	if result != expection {
		t.Errorf("eeeeee")
	}
}
