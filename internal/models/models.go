package models

// Room represents a single room in the ant farm
type Room struct {
	Name        string
	X           int
	Y           int
	IsStart     bool
	IsEnd       bool
	Connections []*Room
}

// AntFarm represents the entire ant colony and maze structure
type AntFarm struct {
	AntCount  int
	Rooms     map[string]*Room
	StartRoom *Room
	EndRoom   *Room
	Input     []string // Stores the original input for output reproduction
}

// Ant represents a single ant in the simulation
type Ant struct {
	ID       int
	Position *Room
	Path     []*Room
}

// Path represents a series of rooms that form a valid path from start to end
type Path struct {
	Rooms []*Room
	Length int
}