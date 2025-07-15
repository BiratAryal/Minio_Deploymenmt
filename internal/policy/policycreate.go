package policy

import "fmt"

func ReadOnly() {
	fmt.Println("Readonly file Creation")
}

func ReadWrite() {
	fmt.Println("Read Write File Creation")
}
func ReadWriteDelete() {
	fmt.Println("Read Write Delete File Creation")
}
func AssignPolicy() {
	fmt.Println("Assign Policy to user")
}
