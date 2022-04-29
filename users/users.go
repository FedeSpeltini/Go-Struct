package users

import "fmt"

type User struct {
	Id      int
	Name    string
	Age     int
	Friends []User
}

func (u User) SayHello() {
	fmt.Printf("Hello, my name is %s.\n", u.Name)
}

func (u *User) AddFriend(friend User) {
	u.Friends = append(u.Friends, friend)
}

func (u User) ListFriends() {
	if len(u.Friends) != 0 {
		fmt.Printf("My friends are: \n")
		for i, friend := range u.Friends {
			fmt.Printf("\t Friend %d: %s\n", i, friend.Name)
		}
	} else {
		fmt.Printf("I don't have any friends.\n")
	}
}

func (u User) FindFriend(name string) (User, bool) {
	for _, friend := range u.Friends {
		if friend.Name == name {
			return friend, true
		}
	}
	return User{}, false
}

func (u *User) RemoveFriend(id int) {
	for i, friend := range u.Friends {
		if friend.Id == id {
			u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
			return
		}
	}
}
