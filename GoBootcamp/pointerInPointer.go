/*
http://www.golangbootcamp.com/book/types#cid20

Looking at the User / Player example, you might have noticed that we composed Player using User but it might be better to compose it with a pointer to a User struct. The reason why a pointer might be better is because in Go, arguments are passed by value and not reference. If you have a small struct that is inexpensive to copy, that is fine, but more than likely, in real life, our User struct will be bigger and should not be copied. Instead we would want to pass by reference (using a pointer). (Section 2.9 & Section 6.3 discuss more in depth how calling a method on a type value vs a pointer affects mutability and memory allocation)

Modify the code to use a pointer but still be able to initialize without using the dot notation.
*/

package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func NewPlayer(Id int, Name, Location string, GameID int) *Player {
	return &Player{
		&User{Id, Name, Location},
		GameID,
	}
}

func main() {
	p := NewPlayer(43, "Matt", "LA", 90404)
	fmt.Println(p.Greetings())
}
