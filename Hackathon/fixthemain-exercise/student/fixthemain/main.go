package main

import "github.com/01-edu/z01"

// TODO: Fix all the errors in this program

// Missing: Define OPEN and CLOSE constants

// Missing: Define Door struct with state field

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	// Missing: Add newline
}

// Missing: Define OpenDoor function

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE // Fixed: was "state = CLOSE"
}

func IsDoorOpen(ptrDoor *Door) bool { // Fixed: parameter type
	PrintStr("is the Door opened ?")
	return ptrDoor.state == OPEN // Fixed: was "Door.state = OPEN"
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE // Fixed: added return statement
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}