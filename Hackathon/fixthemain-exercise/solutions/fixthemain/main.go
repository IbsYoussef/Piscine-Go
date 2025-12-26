package main

import "github.com/01-edu/z01"

// Constants for door state
const (
	OPEN  = true
	CLOSE = false
)

// Door struct with state field
type Door struct {
	state bool
}

// PrintStr prints a string followed by newline
func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// OpenDoor opens the door (sets state to OPEN)
func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
}

// CloseDoor closes the door (sets state to CLOSE)
func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
}

// IsDoorOpen checks if door is open
func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	return ptrDoor.state == OPEN
}

// IsDoorClose checks if door is closed
func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
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
