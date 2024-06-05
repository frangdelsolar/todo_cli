package ui

import "fmt"

type UI struct {}

func (u *UI) Print(message string) {
	fmt.Println(message)
}

