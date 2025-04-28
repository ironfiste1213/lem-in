package algo

import (
	
	"mimo/internal/models" // adjust this to your real import path
)

// FindAllPaths finds all possible paths from start to end in the AntFarm
func FindAllPaths(farm *models.AntFarm) []*models.Path {
	var result []*models.Path

	// Queue of paths (each path is a slice of *Room)
	queue := [][]*models.Room{{farm.StartRoom}}

	for len(queue) > 0 {
		// Dequeue
		currentPath := queue[0]
		queue = queue[1:]

		lastRoom := currentPath[len(currentPath)-1]

		// If we reached the EndRoom, save the path
		if lastRoom == farm.EndRoom {
			newPath := &models.Path{
				Rooms:  currentPath,
				Length: len(currentPath),
			}
			result = append(result, newPath)
			continue
		}

		// Expand: iterate over connected rooms
		for _, neighbor := range lastRoom.Connections {
			// Avoid cycles: don't revisit a room already in path
			if !roomInPath(currentPath, neighbor) {
				// Copy current path and add neighbor
				newPath := append([]*models.Room{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	return result
}

// Helper: check if a room is already in the path
func roomInPath(path []*models.Room, room *models.Room) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}
