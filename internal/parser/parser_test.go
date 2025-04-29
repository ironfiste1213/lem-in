package parser

import (
	"strings"
	"testing"


)

func TestParseSimpleFarm(t *testing.T) {
	input := `
4
##start
A 0 0
B 1 0
##end
C 2 0
A-B
B-C
`

	// Parse the input using your Parse() function
	farm, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check ant count
	if farm.AntCount != 4 {
		t.Errorf("Expected 4 ants, got %d", farm.AntCount)
	}

	// Check start room
	if farm.StartRoom == nil || farm.StartRoom.Name != "A" {
		t.Errorf("Expected start room A, got %v", farm.StartRoom)
	}

	// Check end room
	if farm.EndRoom == nil || farm.EndRoom.Name != "C" {
		t.Errorf("Expected end room C, got %v", farm.EndRoom)
	}

	// Check that rooms are created
	if len(farm.Rooms) != 3 {
		t.Errorf("Expected 3 rooms, got %d", len(farm.Rooms))
	}

	// Check connections
	roomA := farm.Rooms["A"]
	roomB := farm.Rooms["B"]
	roomC := farm.Rooms["C"]

	if len(roomA.Connections) != 1 || roomA.Connections[0].Name != "B" {
		t.Errorf("Room A should be connected to B")
	}
	if len(roomB.Connections) != 2 {
		t.Errorf("Room B should be connected to A and C")
	}
	if len(roomC.Connections) != 1 || roomC.Connections[0].Name != "B" {
		t.Errorf("Room C should be connected to B")
	}
}

func TestParseInvalidFarm(t *testing.T) {
	input := `
3
A 0 0
B 1 0
A-B
`

	_, err := Parse(strings.NewReader(input))
	if err == nil {
		t.Fatalf("Expected an error due to missing start and end rooms, but got nil")
	}
}
