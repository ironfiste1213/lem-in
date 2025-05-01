package algo

import (
	"mimo/internal/models" // adjust this to your real import path
	"sort"
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

// Check if two paths are disjoint (no common rooms except start and end)
func arePathsDisjoint(p1, p2 *models.Path) bool {
	roomSet := make(map[string]bool)
	for i, room := range p1.Rooms {
		if i != 0 && i != len(p1.Rooms)-1 { // exclude start and end
			roomSet[room.Name] = true
		}
	}

	for i, room := range p2.Rooms {
		if i != 0 && i != len(p2.Rooms)-1 { // exclude start and end
			if roomSet[room.Name] {
				return false
			}
		}
	}
	return true
}

// Build all valid groups of disjoint paths
func buildDisjointGroups(paths []*models.Path) [][]*models.Path {
	var groups [][]*models.Path

	var dfs func(current []*models.Path, index int)
	dfs = func(current []*models.Path, index int) {
		if len(current) > 0 {
			temp := make([]*models.Path, len(current))
			copy(temp, current)
			groups = append(groups, temp)
		}
		for i := index; i < len(paths); i++ {
			valid := true
			for _, selected := range current {
				if !arePathsDisjoint(selected, paths[i]) {
					valid = false
					break
				}
			}
			if valid {
				dfs(append(current, paths[i]), i+1)
			}
		}
	}

	dfs([]*models.Path{}, 0)

	return groups
}

// Calculate how many turns needed to move all ants through a group of paths
func calculateTurns(antCount int, paths []*models.Path) int {
	if len(paths) == 0 {
		return 1e9 // big number
	}

	// Paths should be sorted by length ascending (shorter first)
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].Length < paths[j].Length
	})

	turns := 0
	extraAnts := antCount

	// While ants are not all assigned
	for extraAnts > 0 {
		turns++
		for range paths {
			if extraAnts > 0 {
				extraAnts--
			}
		}
	}

	// turns += (length of longest path - 1)
	return turns + paths[len(paths)-1].Length - 1
}

func FindBestGroup(antCount int, allPaths []*models.Path) []*models.Path {
	groups := buildDisjointGroups(allPaths)

	var bestGroup []*models.Path
	bestTurns := 1e9

	for _, group := range groups {
		turns := calculateTurns(antCount, group)
		if float64(turns) < bestTurns {
			bestTurns = float64(turns)
			bestGroup = group
		}
	}

	return bestGroup
}

// Find shortest path using BFS
func findShortestPath(farm *models.AntFarm) []*models.Room {
	type state struct {
		Room *models.Room
		Path []*models.Room
	}
	visited := make(map[string]bool)
	queue := []state{{farm.StartRoom, []*models.Room{farm.StartRoom}}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Room == farm.EndRoom {
			return current.Path
		}

		for _, neighbor := range current.Room.Connections {
			if visited[neighbor.Name] || roomInPath(current.Path, neighbor) {
				continue
			}
			visited[neighbor.Name] = true
			newPath := append([]*models.Room{}, current.Path...)
			newPath = append(newPath, neighbor)
			queue = append(queue, state{neighbor, newPath})
		}
	}
	return nil
}

// Find up to maxPaths disjoint paths
func FindKDisjointPaths(farm *models.AntFarm, maxPaths int) []*models.Path {
	var result []*models.Path

	for i := 0; i < maxPaths; i++ {
		path := findShortestPath(farm)
		if path == nil {
			break
		}
		// Remove intermediate rooms from graph to ensure disjoint paths
		for _, room := range path[1 : len(path)-1] {
			delete(farm.Rooms, room.Name)
		}
		result = append(result, &models.Path{Rooms: path, Length: len(path)})
	}

	return result
}
