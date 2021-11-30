package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	ret := fmt.Sprintf("Welcome to my party, %s!", name)
	return ret
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	ret := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return ret
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	ret := fmt.Sprintf("Welcome to my party, %s!\n", name)
	ret += fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n", table, direction, distance)
	ret += fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return ret
}
