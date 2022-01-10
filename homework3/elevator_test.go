package main

import "testing"

/**
go test -v elevator_test.go elevator.go
*/
func TestElevator_LetsGo(t *testing.T) {
	e := &Elevator{
		CurrentFloor: 1,
		Direction:    "up",
	}
	path := e.LetsGo()
	if len(path) != 0 {
		t.Fatalf("电梯不应该动...")
	}
}
func TestElevator_LetsGo2(t *testing.T) {
	e := &Elevator{
		CurrentFloor: 1,
		Direction:    "up",
		Destinations: []int{3},
	}
	path := e.LetsGo()
	if len(path) != 1 {
		t.Fatalf("电梯行动轨迹不对")
	}
	pathExpected := []int{3}
	for k, v := range path {
		if v != pathExpected[k] {
			t.Fatalf("电梯轨迹不对，应该是%v，实际是%v", pathExpected, path)
		}
	}
}
func TestElevator_LetsGo3(t *testing.T) {
	e := &Elevator{
		CurrentFloor: 3,
		Direction:    "up",
		Destinations: []int{2, 4},
	}
	path := e.LetsGo()
	if len(path) != 2 {
		t.Fatalf("电梯行动轨迹不对")
	}
	pathExpected := []int{4, 2}
	for k, v := range path {
		if v != pathExpected[k] {
			t.Fatalf("电梯轨迹不对，应该是%v，实际是%v", pathExpected, path)
		}
	}
}

func TestElevator_LetsGo4(t *testing.T) {
	e := &Elevator{
		CurrentFloor: 3,
		Direction:    "up",
		Destinations: []int{2, 4, 5},
	}
	path := e.LetsGo()
	if len(path) != 3 {
		t.Fatalf("电梯行动轨迹不对")
	}
	pathExpected := []int{4, 5, 2}
	for k, v := range path {
		if v != pathExpected[k] {
			t.Fatalf("电梯轨迹不对，应该是%v，实际是%v", pathExpected, path)
		}
	}
}
