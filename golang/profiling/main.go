package main

import "fmt"

type User struct {
	name string
}

var users []*User

func main() {
	u := createUser() // u points to heap memory address
	fmt.Println(u)

	cu := createCopyUser()
	fmt.Println(cu)

	user1 := &User{}
	doesNotEscape(user1)

	user2 := &User{}
	doesEscape(user2)

	user3 := &User{}
	doesEscapeV2(user3)
}

func createUser() *User {
	user := &User{"yerken"} //this is allocated in the heap,
	// because escape analysis determined variable pointer is returned from the function down below
	return user
}

func createCopyUser() User {
	user := &User{"yerken"} //this is allocated in the stack,
	return *user
}

func doesNotEscape(user *User) {
	user.name = ""
}

func doesEscape(user *User) {
	users = append(users, user)
}

func doesEscapeV2(user *User) *User {
	user.name = "123"
	doesEscape(user)
	return user
}
