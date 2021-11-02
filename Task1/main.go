package main

import "fmt"

type User struct {
	name   string
	age    int
	merrid bool
}

type UserAuthorization struct {
	login                  string
	password               string
	numberOfwrongPasswords int
	isAuthorithed          bool
	User
}

type UsersAuth []UserAuthorization

func main() {
	array := UsersInput()
	array.logic(UsersInput())
}

//making all users ...

func UsersInput() UsersAuth {
	allUsers := UsersAuth{
		UserAuthorization{login: "Bryan", password: "Human", numberOfwrongPasswords: 1, isAuthorithed: true,
			User: User{name: "Bryan", age: 45, merrid: true}},

		UserAuthorization{login: "Jane", password: "Rocker", numberOfwrongPasswords: 2, isAuthorithed: true,
			User: User{name: "Jane", age: 25, merrid: false}},

		UserAuthorization{login: "Nancy", password: "Mother", numberOfwrongPasswords: 3, isAuthorithed: false,
			User: User{name: "Nancy", age: 41, merrid: true}},

		UserAuthorization{login: "Chris", password: "Dude", numberOfwrongPasswords: 1, isAuthorithed: false,
			User: User{name: "Chris", age: 34, merrid: true}},

		UserAuthorization{login: "Martha", password: "Cook", numberOfwrongPasswords: 2, isAuthorithed: true,
			User: User{name: "Martha", age: 55, merrid: true}},
	}

	//another mathod of adding a component...

	usr := User{name: "Jone", age: 21, merrid: false}
	usrAuth := UserAuthorization{"Jone", "Markis", 2, false, usr}
	allUsers = append(allUsers, usrAuth)
	return allUsers
}

// las part of the task printing the structer of values

func (usrAuth UsersAuth) logic(ArrayOfValues UsersAuth) {
	for i, value := range ArrayOfValues {
		outString := fmt.Sprintf("<%[1]d> value: %[2]v", i, value)
		fmt.Println(outString)
		
	}

}
