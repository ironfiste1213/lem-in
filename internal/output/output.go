package output

import (
	"fmt"
	"mimo/internal/models"
	"sort"
)

type SimAnt struct {
	ID       int
	Path     []*models.Room
	Step     int
	Finished bool
}

// Assign ants to paths smartly
func assignAntsSmart(farm *models.AntFarm, paths []*models.Path) []*SimAnt {
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].Length < paths[j].Length
	})

	// How many ants per path
	antsPerPath := make([]int, len(paths))
	remainingAnts := farm.AntCount

	// Smart assignment
	for {
		progress := false
		for i := 0; i < len(paths); i++ {
			if remainingAnts > 0 {
				antsPerPath[i]++
				remainingAnts--
				progress = true
			}
		}
		if !progress {
			break
		}
	}

	// Create ants with assigned paths
	ants := make([]*SimAnt, 0, farm.AntCount)
	antID := 1
	for pathIdx, count := range antsPerPath {
		for j := 0; j < count; j++ {
			ant := &SimAnt{
				ID:   antID,
				Path: paths[pathIdx].Rooms,
				Step: 0,
			}
			ants = append(ants, ant)
			antID++
		}
	}

	return ants
}

// Simulate ants moving along the selected paths
func SimulateAntsSmart(farm *models.AntFarm, paths []*models.Path) {
	if len(paths) == 0 {
		fmt.Println("No paths available!")
		return
	}

	ants := assignAntsSmart(farm, paths)

	movingAnts := ants
	for {
		moveLine := ""

		for _, ant := range movingAnts {
			if !ant.Finished && ant.Step < len(ant.Path)-1 {
				ant.Step++
				moveLine += fmt.Sprintf("L%d-%s ", ant.ID, ant.Path[ant.Step].Name)
				if ant.Path[ant.Step] == farm.EndRoom {
					ant.Finished = true
				}
			}
		}

		if moveLine != "" {
			fmt.Println(moveLine)
		}

		allFinished := true
		for _, ant := range movingAnts {
			if !ant.Finished {
				allFinished = false
				break
			}
		}
		if allFinished {
			break
		}
	}
}
