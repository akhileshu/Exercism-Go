package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %v!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %v! You are now %v years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	formatedTable := fmt.Sprintf("%03d", table)       // 0 = pad with zeros, 3 = minimum width, d = decimal integer.
	formatedDistance := fmt.Sprintf("%.1f", distance) // %.1f means: format a floating-point number with 1 digit after the decimal point.
	return fmt.Sprintf("Welcome to my party, %v!\nYou have been assigned to table %v. Your table is %v, exactly %v meters from here.\nYou will be sitting next to %v.", name, formatedTable, direction, formatedDistance, neighbor)
}
