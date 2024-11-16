// 1-  The Single Responsibility Principle

// The Single Responsibility Principle states that a class or module should have only one reason to change.
// In other words, a type should have a single responsibility, making the code easier to understand and maintain.

package main

	FName string
	LName string
}

type Users map[int]*User

// Dab Practice
// Because User has multiple responsibilities getfull name and save it
func (u *User) GetFullName() string {
	return u.FName + " " + u.LName
}
func (u *User) Save() {
	Users[u.Id] = u
}

//Good Practice is to separite these responsibilities

type UserRepository struct {
}

func (ur *UserRepository) Save(u *User) {
	Users[u.Id] = u
}

func main() {

}
