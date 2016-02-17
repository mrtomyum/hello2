package main

import (
	"fmt"
)

type State int

func (s *State) On() {
	*s = ON //เนื่องจาก s *State ... s เป็นป้ายบอกที่อยู่ เก็บค่าไม่ได้ การจะรับค่าจึงต้องใช้ *s แทน
}

func (s *State) Off() {
	*s = OFF
}

func (s State) ShowState() string {
	switch s {
	case ON:
		return "ON"
	case OFF:
		return "OFF"
	}
	return ""
}

func (s State) String() string {
	switch s {
	case ON:
		return "ON"
	case OFF:
		return "OFF"
	}
	return ""
}

const (
	OFF State = 0
	ON  State = 1
)

var currentState State

func Turn(s State) {
	currentState = s
}

func ShowState(s State) string {
	switch s {
	case ON:
		return "ON"
	case OFF:
		return "OFF"
	}
	return ""
}

func main() {
	currentState.On()
	fmt.Println(currentState)
	currentState.Off()
	fmt.Println(ShowState(currentState))
}
