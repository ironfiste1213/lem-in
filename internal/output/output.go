package output

import (
	"fmt"
	"mimo/internal/models"
	"sort"
	"strings"
)

type SimAnt struct {
	ID       int
	Path     []*models.Room
	Step     int
	Finished bool
}

func assignAntsSmart(farm *models.AntFarm, paths []*models.Path) []*SimAnt {
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].Length < paths[j].Length
	})

	antsPerPath := make([]int, len(paths))
	remainingAnts := farm.AntCount

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

func SimulateAntsSmart(farm *models.AntFarm, paths []*models.Path) {
	if len(paths) == 0 {
		fmt.Println("No paths available!")
		return
	}

	ants := assignAntsSmart(farm, paths)

	movingAnts := []*SimAnt{}
	waitingAnts := make([]*SimAnt, len(ants))
	copy(waitingAnts, ants)

	for {
		moveLine := ""
		occupied := make(map[string]bool)

		// Mark rooms currently occupied (not Start or End)
		for _, ant := range movingAnts {
			if !ant.Finished && ant.Step > 0 && ant.Step < len(ant.Path)-1 {
				occupied[ant.Path[ant.Step].Name] = true
			}
		}

		newMovingAnts := []*SimAnt{}

		// 🚀 First move ants already on paths
		for _, ant := range movingAnts {
			if ant.Finished {
				continue
			}

			nextRoom := ant.Path[ant.Step+1]
			if nextRoom.IsEnd || !occupied[nextRoom.Name] {
				ant.Step++
				moveLine += fmt.Sprintf("L%d-%s ", ant.ID, nextRoom.Name)
				if nextRoom.IsEnd {
					ant.Finished = true
				} else {
					occupied[nextRoom.Name] = true
					newMovingAnts = append(newMovingAnts, ant)
				}
			} else {
				newMovingAnts = append(newMovingAnts, ant)
			}
		}

		// 🚀 Then start NEW ants if possible (after moving)
		for len(waitingAnts) > 0 {
			nextRoom := waitingAnts[0].Path[1]
			if !occupied[nextRoom.Name] || nextRoom.IsEnd {
				ant := waitingAnts[0]
				waitingAnts = waitingAnts[1:]

				ant.Step = 1
				moveLine += fmt.Sprintf("L%d-%s ", ant.ID, nextRoom.Name)
				if nextRoom.IsEnd {
					ant.Finished = true
				} else {
					occupied[nextRoom.Name] = true
					newMovingAnts = append(newMovingAnts, ant)
				}
			} else {
				break // cannot start more ants now
			}
		}

		movingAnts = newMovingAnts

		if moveLine != "" {
			fmt.Println(strings.TrimSpace(moveLine))
		}

		if len(movingAnts) == 0 && len(waitingAnts) == 0 {
			break
		}
	}
}
