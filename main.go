package main

import "github.com/my-first-struct/users"

func main() {

	fede := users.User{Id: 1, Name: "Fede", Age: 30}

	jose := users.User{Id: 2, Name: "Jose", Age: 25}
	ana := users.User{Id: 3, Name: "Ana", Age: 25}

	fede.SayHello()
	fede.AddFriend(jose)
	fede.AddFriend(ana)
	fede.ListFriends()
}
